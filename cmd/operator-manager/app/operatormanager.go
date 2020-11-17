package app

import (
	"github.com/wflysnow/my-middleware-operator/cmd/operator-manager/app/options"
	redisInformerFactory "github.com/wflysnow/my-middleware-operator/pkg/clients/informers/externalversions"
	"github.com/wflysnow/my-middleware-operator/pkg/operator"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/controller"
)

const (
	OperatorStartJitter = 1.0
)

type OperatorContext struct {
	// ClientBuilder will provide a client for this operator to use
	OperatorClientBuilder operator.OperatorClientBuilder

	kubeConfig *restclient.Config

	// ClientBuilder will provide a default client for this operator to use
	DefaultClientBuilder controller.ControllerClientBuilder

	// RedisInformerFactory gives access to informers for the operator.
	RedisInformerFactory redisInformerFactory.SharedInformerFactory

	// InformerFactory gives access to informers for the operator.
	InformerFactory informers.SharedInformerFactory

	// Options provides access to init options for a given operator
	Options options.OperatorManagerServer

	// AvailableResources is a map listing currently available resources
	//AvailableResources map[schema.GroupVersionResource]bool

	// Stop is the stop channel
	Stop <-chan struct{}

	// InformersStarted is closed after all of the controllers have been initialized and are running.  After this point it is safe,
	// for an individual controller to start the shared informers. Before it is closed, they should not.
	InformersStarted chan struct{}
}

func KnownOperators() []string {
	res := sets.StringKeySet(NewOperatorInitializers())
	return res.List()

}

type InitFunc func(OperatorContext) (bool, error)

func NewOperatorInitializers() map[string]InitFunc {
	controllers := map[string]InitFunc{}
	controllers["rediscluster"] = startRedisClusterController
	return controllers
}

func Run(o *options.OperatorManagerServer) error {
	return nil
}

func startRedisClusterController(otx OperatorContext) (bool, error) {

	return false, nil
}
