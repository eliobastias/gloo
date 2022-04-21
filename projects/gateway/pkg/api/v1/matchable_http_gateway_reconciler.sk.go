// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionMatchableHttpGatewayFunc func(original, desired *MatchableHttpGateway) (bool, error)

type MatchableHttpGatewayReconciler interface {
	Reconcile(namespace string, desiredResources MatchableHttpGatewayList, transition TransitionMatchableHttpGatewayFunc, opts clients.ListOpts) error
}

func matchableHttpGatewaysToResources(list MatchableHttpGatewayList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, matchableHttpGateway := range list {
		resourceList = append(resourceList, matchableHttpGateway)
	}
	return resourceList
}

func NewMatchableHttpGatewayReconciler(client MatchableHttpGatewayClient, statusSetter resources.StatusSetter) MatchableHttpGatewayReconciler {
	return &matchableHttpGatewayReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type matchableHttpGatewayReconciler struct {
	base reconcile.Reconciler
}

func (r *matchableHttpGatewayReconciler) Reconcile(namespace string, desiredResources MatchableHttpGatewayList, transition TransitionMatchableHttpGatewayFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "matchableHttpGateway_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*MatchableHttpGateway), desired.(*MatchableHttpGateway))
		}
	}
	return r.base.Reconcile(namespace, matchableHttpGatewaysToResources(desiredResources), transitionResources, opts)
}