package app

import (
	"fmt"
	"os"
	"time"

	"k8s.io/client-go/tools/leaderelection/resourcelock"

	"github.com/golang/glog"
	"github.com/wflysnow/my-middleware-operator/cmd/operator-manager/app/options"
	redisInformerFactory "github.com/wflysnow/my-middleware-operator/pkg/clients/informers/externalversions"
	"github.com/wflysnow/my-middleware-operator/pkg/operator"
	extensionclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/version"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/record"
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

// Run run the OMserver. This should never exit
func Run(s *options.OperatorManagerServer) error {
	glog.Infof("Version: %+v", version.Get())

	kubeClient, leaderElectionClient, extensionCRClient, kubeconfig, err := createClients(s)
	if err != nil {
		return err
	}
	err = CreateRedisClusterCRD(extensionCRClient)
	if err != nil {
		if errors.IsAlreadyExists(err) {
			glog.Infof("redis cluster crd is already craeted.")
		} else {
			fmt.Fprint(os.Stderr, err)
			return err
		}
	}

	go startHTTP(s)

	recorder := createRecorder(kubeClient)
	run := func(stop <-chan struct{}) {
		operatorClientBuilder := operator.SimpleOperatorClientBuilder{
			ClientConfig: kubeconfig,
		}

		rootClientBuilder := controller.SimpleControllerClientBuilder{
			ClientConfig: kubeconfig,
		}
		otx, err := CreateOperatorContext(s, kubeconfig, operatorClientBuilder, rootClientBuilder, stop)
		if err != nil {
			glog.Fatalf("error building controller context: %v", err)
		}
		otx.InformerFactory = informers.NewSharedInformerFactory(kubeClient, time.Duration(s.ResyncPeriod)*time.Second)

		if err := StartOperators(); err != nil {
			glog.Fatalf("error starting operators: %v", err)
		}

		otx.RedisInformerFactory.Start(otx.Stop)
		otx.InformerFactory.Start(otx.Stop)
		close(otx.InformersStarted)
		select {}
	}
	if !s.LeaderElection.LeaderElect {
		run(nil)
		panic("unreachable")
	}
	id, err := os.Hostname()
	if err != nil {
		return err
	}
	rl, err := resourcelock.New(s.LeaderElection.ResourceLock,
		"kube-system",
		"middleware-operator-manager",
		leaderElectionClient.CoreV1(),
		resourcelock.ResourceLockConfig{
			Identity:      id,
			EventRecorder: recorder,
		})
	if err != nil {
		glog.Fatalf("error creating lock: %v", err)
	}
	leaderelection.RunOrDie(leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: s.LeaderElection.LeaseDuration.Duration,
		RenewDeadline: s.LeaderElection.RenewDeadline.Duration,
		RetryPeriod:   s.LeaderElection.RetryPeriod.Duration,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				glog.Fatalf("leaderelection lost")
			},
		},
	})
	panic("unreachable")
}

func createClients(s *options.OperatorManagerServer) (*clientset.Clientset, *clientset.Clientset, *extensionclient.Clientset, *restclient.Config, error) {
	return nil, nil, nil, &restclient.Config{}, nil
}

func CreateRedisClusterCRD(cextensionCRClient *extensionclient.Clientset) error {
	return nil
}

func startHTTP(s *options.OperatorManagerServer) {

}

func createRecorder(kubeClient *clientset.Clientset) record.EventRecorder {
	return nil
}

func CreateOperatorContext(s *options.OperatorManagerServer, kubeconfig *restclient.Config, operatorClientBuilder operator.OperatorClientBuilder, rootClientBuilder controller.ControllerClientBuilder, stop <-chan struct{}) (OperatorContext, error) {
	return OperatorContext{}, nil
}

func StartOperators() error {
	return nil
}
