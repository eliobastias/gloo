package httproute

import (
	"container/list"
	"context"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gateway2/parameters"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	"github.com/solo-io/gloo/projects/gateway2/translator/backendref"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/registry"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
)

var (
	awsMissingFuncRefError                = eris.New("upstreams must have a logical name specified in the backend ref via the parameters extensionref")
	azureMissingFuncRefError              = eris.New("upstreams must have a function name specified in the backend ref via the parameters extensionref")
	nonFunctionUpstreamWithParameterError = eris.New("parameters extensionref is only supported for aws and azure upstreams")
)

func TranslateGatewayHTTPRouteRules(
	ctx context.Context,
	pluginRegistry registry.PluginRegistry,
	queries query.GatewayQueries,
	gwListener gwv1.Listener,
	route gwv1.HTTPRoute,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
) []*v1.Route {
	var finalRoutes []*v1.Route
	routesVisited := sets.New[types.NamespacedName]()

	// Hostnames need to be explicitly passed to the plugins since they
	// are required by delegatee (child) routes of delegated routes that
	// won't have spec.Hostnames set.
	hostnames := make([]gwv1.Hostname, len(route.Spec.Hostnames))
	copy(hostnames, route.Spec.Hostnames)

	delegationChain := list.New()

	translateGatewayHTTPRouteRulesUtil(
		ctx, pluginRegistry, queries, gwListener, route, reporter, baseReporter, &finalRoutes, routesVisited, hostnames, delegationChain)
	return finalRoutes
}

// translateGatewayHTTPRouteRulesUtil is a helper to translate an HTTPRoute.
// In case of route delegation, this function is recursively invoked to flatten the delegated route tree.
func translateGatewayHTTPRouteRulesUtil(
	ctx context.Context,
	pluginRegistry registry.PluginRegistry,
	queries query.GatewayQueries,
	gwListener gwv1.Listener,
	route gwv1.HTTPRoute,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	outputs *[]*v1.Route,
	routesVisited sets.Set[types.NamespacedName],
	hostnames []gwv1.Hostname,
	delegationChain *list.List,
) {
	for _, rule := range route.Spec.Rules {
		rule := rule
		if rule.Matches == nil {
			// from the spec:
			// If no matches are specified, the default is a prefix path match on “/”, which has the effect of matching every HTTP request.
			rule.Matches = []gwv1.HTTPRouteMatch{{}}
		}

		outputRoutes := translateGatewayHTTPRouteRule(
			ctx,
			pluginRegistry,
			queries,
			gwListener,
			&route,
			rule,
			reporter,
			baseReporter,
			outputs,
			routesVisited,
			hostnames,
			delegationChain,
		)
		for _, outputRoute := range outputRoutes {
			// The above function will return a nil route if a matcher fails to apply plugins
			// properly. This is a signal to the caller that the route should be dropped.
			if outputRoute == nil {
				continue
			}

			*outputs = append(*outputs, outputRoute)
		}
	}
}

func translateGatewayHTTPRouteRule(
	ctx context.Context,
	pluginRegistry registry.PluginRegistry,
	queries query.GatewayQueries,
	gwListener gwv1.Listener,
	gwroute *gwv1.HTTPRoute,
	rule gwv1.HTTPRouteRule,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	outputs *[]*v1.Route,
	routesVisited sets.Set[types.NamespacedName],
	hostnames []gwv1.Hostname,
	delegationChain *list.List,
) []*v1.Route {
	routes := make([]*v1.Route, len(rule.Matches))
	for idx, match := range rule.Matches {
		match := match // pike
		outputRoute := &v1.Route{
			Matchers: []*matchers.Matcher{translateGlooMatcher(match)},
			Action:   nil,
			Options:  &v1.RouteOptions{},
		}

		var delegatedRoutes []*v1.Route
		var delegates bool
		if len(rule.BackendRefs) > 0 {
			delegates = setRouteAction(
				ctx,
				queries,
				gwroute,
				rule,
				outputRoute,
				reporter,
				baseReporter,
				pluginRegistry,
				gwListener,
				match,
				&delegatedRoutes,
				routesVisited,
				delegationChain,
			)
		}

		rtCtx := &plugins.RouteContext{
			Listener:        &gwListener,
			Route:           gwroute,
			Hostnames:       hostnames,
			DelegationChain: delegationChain,
			Rule:            &rule,
			Match:           &match,
			Reporter:        reporter,
		}

		// Apply the plugins for this route
		for _, plugin := range pluginRegistry.GetRoutePlugins() {
			err := plugin.ApplyRoutePlugin(ctx, rtCtx, outputRoute)
			if err != nil {
				contextutils.LoggerFrom(ctx).Errorf("error in RoutePlugin: %v", err)
			}

			// If this parent route has delegatee routes, override any applied policies
			// that are on the child with the parent's policies.
			// When a plugin is invoked on a route, it must override the existing route.
			for _, child := range delegatedRoutes {
				err := plugin.ApplyRoutePlugin(ctx, rtCtx, child)
				if err != nil {
					contextutils.LoggerFrom(ctx).Errorf("error applying RoutePlugin to child route %s: %v", child.GetName(), err)
				}
			}
		}
		// Add the delegatee output routes to the final output list
		*outputs = append(*outputs, delegatedRoutes...)

		// It is possible for a parent route to not produce an output route action
		// if it only delegates and does not directly route to a backend.
		// We should only set a direct response action when there is no output action
		// for a parent rule and when there are no delegated routes because this would
		// otherwise result in a top level matcher with a direct response action for the
		// path that the parent is delegating for.
		if outputRoute.GetAction() == nil && !delegates {
			outputRoute.Action = &v1.Route_DirectResponseAction{
				DirectResponseAction: &v1.DirectResponseAction{
					Status: http.StatusInternalServerError,
				},
			}
		}

		// A parent route that delegates to a child route should not have an output route
		// action (outputRoute.Action) as the routes are derived from the child route.
		// So this conditional ensures that we do not create a top level route matcher
		// for the parent route when it delegates to a child route.
		if outputRoute.GetAction() != nil {
			outputRoute.Matchers = []*matchers.Matcher{translateGlooMatcher(match)}
			routes[idx] = outputRoute
		}
	}
	return routes
}

func translateGlooMatcher(match gwv1.HTTPRouteMatch) *matchers.Matcher {
	// headers
	headers := make([]*matchers.HeaderMatcher, 0, len(match.Headers))
	for _, header := range match.Headers {
		h := translateGlooHeaderMatcher(header)
		if h != nil {
			headers = append(headers, h)
		}
	}

	// query params
	var queryParamMatchers []*matchers.QueryParameterMatcher
	for _, param := range match.QueryParams {
		queryParamMatchers = append(queryParamMatchers, &matchers.QueryParameterMatcher{
			Name:  string(param.Name),
			Value: param.Value,
			Regex: false,
		})
	}

	// set path
	pathType, pathValue := parsePath(match.Path)

	var methods []string
	if match.Method != nil {
		methods = []string{string(*match.Method)}
	}
	m := &matchers.Matcher{
		// CaseSensitive:   nil,
		Headers:         headers,
		QueryParameters: queryParamMatchers,
		Methods:         methods,
	}

	switch pathType {
	case gwv1.PathMatchPathPrefix:
		m.PathSpecifier = &matchers.Matcher_Prefix{
			Prefix: pathValue,
		}
	case gwv1.PathMatchExact:
		m.PathSpecifier = &matchers.Matcher_Exact{
			Exact: pathValue,
		}
	case gwv1.PathMatchRegularExpression:
		m.PathSpecifier = &matchers.Matcher_Regex{
			Regex: pathValue,
		}
	}

	return m
}

func translateGlooHeaderMatcher(header gwv1.HTTPHeaderMatch) *matchers.HeaderMatcher {
	regex := false
	if header.Type != nil && *header.Type == gwv1.HeaderMatchRegularExpression {
		regex = true
	}

	return &matchers.HeaderMatcher{
		Name:  string(header.Name),
		Value: header.Value,
		Regex: regex,
		// InvertMatch: header.InvertMatch,
	}
}

func parsePath(path *gwv1.HTTPPathMatch) (gwv1.PathMatchType, string) {
	pathType := gwv1.PathMatchPathPrefix
	pathValue := "/"
	if path != nil && path.Type != nil {
		pathType = *path.Type
	}
	if path != nil && path.Value != nil {
		pathValue = *path.Value
	}
	return pathType, pathValue
}

func setRouteAction(
	ctx context.Context,
	queries query.GatewayQueries,
	gwroute *gwv1.HTTPRoute,
	rule gwv1.HTTPRouteRule,
	outputRoute *v1.Route,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	pluginRegistry registry.PluginRegistry,
	gwListener gwv1.Listener,
	match gwv1.HTTPRouteMatch,
	outputs *[]*v1.Route,
	routesVisited sets.Set[types.NamespacedName],
	delegationChain *list.List,
) bool {
	var weightedDestinations []*v1.WeightedDestination
	backendRefs := rule.BackendRefs
	delegates := false

	for _, backendRef := range backendRefs {
		// If the backend is an HTTPRoute, it implies route delegation
		// for which delegated routes are recursively flattened and translated
		if backendref.RefIsHTTPRoute(backendRef.BackendObjectReference) {
			delegates = true
			// Flatten delegated HTTPRoute references
			err := flattenDelegatedRoutes(
				ctx, queries, gwroute, backendRef, reporter, baseReporter, pluginRegistry, gwListener, match, outputs, routesVisited, delegationChain)
			if err != nil {
				reporter.SetCondition(reports.HTTPRouteCondition{
					Type:    gwv1.RouteConditionResolvedRefs,
					Status:  metav1.ConditionFalse,
					Reason:  gwv1.RouteReasonRefNotPermitted,
					Message: err.Error(),
				})
			}
			continue
		}

		clusterName := "blackhole_cluster"
		ns := "blackhole_ns"

		obj, err := queries.GetBackendForRef(ctx, queries.ObjToFrom(gwroute), &backendRef.BackendObjectReference)
		ptrClusterName := query.ProcessBackendRef(obj, err, reporter, backendRef.BackendObjectReference)
		if ptrClusterName != nil {
			clusterName = *ptrClusterName
			ns = obj.GetNamespace()
		}

		var weight *wrappers.UInt32Value
		if backendRef.Weight != nil {
			weight = &wrappers.UInt32Value{
				Value: uint32(*backendRef.Weight),
			}
		} else {
			// according to spec, default weight is 1
			weight = &wrappers.UInt32Value{
				Value: 1,
			}
		}

		var port uint32
		if backendRef.Port != nil {
			port = uint32(*backendRef.Port)
		}

		switch {
		// get backend for ref - we must do it to make sure we have permissions to access it.
		// also we need the service so we can translate its name correctly.
		case backendref.RefIsService(backendRef.BackendObjectReference):
			weightedDestinations = append(weightedDestinations, &v1.WeightedDestination{
				Destination: &v1.Destination{
					DestinationType: &v1.Destination_Kube{
						Kube: &v1.KubernetesServiceDestination{
							Ref: &core.ResourceRef{
								Name:      clusterName,
								Namespace: ns,
							},
							Port: port,
						},
					},
				},
				Weight:  weight,
				Options: nil,
			})

		case backendref.RefIsUpstream(backendRef.BackendObjectReference):
			upstream, ok := obj.(*gloov1.Upstream)
			if !ok {
				// should never happen
				contextutils.LoggerFrom(ctx).Errorf("expected upstream, got %T", obj)
				continue
			}
			spec, err := makeDestinationSpec(upstream, backendRef.Filters)
			if err != nil {
				reporter.SetCondition(reports.HTTPRouteCondition{
					Type:    gwv1.RouteConditionResolvedRefs,
					Status:  metav1.ConditionFalse,
					Reason:  gwv1.RouteReasonBackendNotFound,
					Message: err.Error(),
				})
				contextutils.LoggerFrom(ctx).Errorf("failed to make destination spec for backend upstream %s: %v", upstream.Name, err)
				continue
			}
			weightedDestinations = append(weightedDestinations, &v1.WeightedDestination{
				Destination: &v1.Destination{
					DestinationType: &v1.Destination_Upstream{
						Upstream: &core.ResourceRef{
							Name:      clusterName,
							Namespace: ns,
						},
					},
					DestinationSpec: spec,
				},
				Weight:  weight,
				Options: nil,
			})

		default:
			contextutils.LoggerFrom(ctx).Errorf("unsupported backend type for kind: %v and type: %v", *backendRef.BackendObjectReference.Kind, *backendRef.BackendObjectReference.Group)
		}
	}

	// TODO(revert): need to add ClusterNotFoundResponseCode: routev3.RouteAction_INTERNAL_SERVER_ERROR,

	switch len(weightedDestinations) {
	case 0:
		// True for delegated BackendRefs. Nothing to be done here since the recursive
		// implementation of creating routes should correctly configure the route action
	case 1:
		outputRoute.Action = &v1.Route_RouteAction{
			RouteAction: &v1.RouteAction{
				Destination: &v1.RouteAction_Single{Single: weightedDestinations[0].GetDestination()},
			},
		}
	default:
		outputRoute.Action = &v1.Route_RouteAction{
			RouteAction: &v1.RouteAction{
				Destination: &v1.RouteAction_Multi{Multi: &v1.MultiDestination{
					Destinations: weightedDestinations,
				}},
			},
		}
	}

	return delegates
}

// makeDestinationSpec computes the destination spec for a given upstream based on the type of upstream and the Filters from the backend reference
func makeDestinationSpec(upstream *gloov1.Upstream, filters []gwv1.HTTPRouteFilter) (*v1.DestinationSpec, error) {
	var sectionName string
	for _, filter := range filters {
		// only look for 'parameters' extensionref filters
		if filter.Type != gwv1.HTTPRouteFilterExtensionRef ||
			filter.ExtensionRef == nil ||
			filter.ExtensionRef.Group != parameters.ParameterGroup ||
			filter.ExtensionRef.Kind != parameters.ParameterKind {
			continue
		}
		sectionName = string(filter.ExtensionRef.Name)
		break
	}

	switch upstream.Spec.GetUpstreamType().(type) {
	case *v1.Upstream_Aws:
		if sectionName == "" {
			return nil, awsMissingFuncRefError
		}
		return &v1.DestinationSpec{
			DestinationType: &v1.DestinationSpec_Aws{
				Aws: &aws.DestinationSpec{
					LogicalName: sectionName,
				},
			},
		}, nil
	case *v1.Upstream_Azure:
		if sectionName == "" {
			return nil, azureMissingFuncRefError
		}
		return &v1.DestinationSpec{
			DestinationType: &v1.DestinationSpec_Azure{
				Azure: &azure.DestinationSpec{
					FunctionName: sectionName,
				},
			},
		}, nil
	}

	// not a supported upstream type
	if sectionName != "" {
		return nil, nonFunctionUpstreamWithParameterError
	}
	return nil, nil
}
