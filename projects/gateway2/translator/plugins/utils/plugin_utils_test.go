package utils_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/gomega"

	sologatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	solokubev1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/utils"
	"github.com/solo-io/gloo/projects/gateway2/translator/testutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/faultinjection"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestExtensionRef(t *testing.T) {
	g := NewWithT(t)
	deps := []client.Object{routeOption()}
	queries := testutils.BuildGatewayQueries(deps)

	rtCtx := routeContext()
	gk := schema.GroupKind{
		Group: sologatewayv1.RouteOptionGVK.Group,
		Kind:  sologatewayv1.RouteOptionGVK.Kind,
	}
	filter := utils.FindExtensionRefFilter(rtCtx.Rule, gk)
	g.Expect(filter).ToNot(BeNil())

	routeOption, err := utils.GetExtensionRefObj[*solokubev1.RouteOption](context.Background(), rtCtx.Route, queries, filter.ExtensionRef)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(routeOption.Spec.GetOptions().GetFaults().GetAbort().GetPercentage()).To(BeEquivalentTo(1))
}

func TestExtensionRefWrongObject(t *testing.T) {
	g := NewWithT(t)
	deps := []client.Object{routeOption()}
	queries := testutils.BuildGatewayQueries(deps)

	rtCtx := routeContext()
	gk := schema.GroupKind{
		Group: sologatewayv1.RouteOptionGVK.Group,
		Kind:  sologatewayv1.RouteOptionGVK.Kind,
	}
	filter := utils.FindExtensionRefFilter(rtCtx.Rule, gk)
	g.Expect(filter).ToNot(BeNil())

	_, err := utils.GetExtensionRefObj[*solokubev1.VirtualHostOption](context.Background(), rtCtx.Route, queries, filter.ExtensionRef)
	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Is(err, utils.ErrTypesNotEqual)).To(BeTrue())
}

func routeOption() *solokubev1.RouteOption {
	return &solokubev1.RouteOption{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "policy",
			Namespace: "default",
		},
		Spec: sologatewayv1.RouteOption{
			Options: &v1.RouteOptions{
				Faults: &faultinjection.RouteFaults{
					Abort: &faultinjection.RouteAbort{
						Percentage: 1.00,
					},
				},
			},
		},
	}
}

func routeContext() plugins.RouteContext {
	return plugins.RouteContext{
		Route: &gwv1.HTTPRoute{},
		Rule: &gwv1.HTTPRouteRule{
			Filters: []gwv1.HTTPRouteFilter{
				{
					Type: gwv1.HTTPRouteFilterExtensionRef,
					ExtensionRef: &gwv1.LocalObjectReference{
						Group: gwv1.Group(sologatewayv1.RouteOptionGVK.Group),
						Kind:  gwv1.Kind(sologatewayv1.RouteOptionGVK.Kind),
						Name:  "policy",
					},
				},
			},
		},
	}
}
