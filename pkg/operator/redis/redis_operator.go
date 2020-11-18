package redis

import (
	"time"

	"github.com/golang/glog"
	"github.com/wflysnow/my-middleware-operator/cmd/operator-manager/app/options"
	redistype "github.com/wflysnow/my-middleware-operator/pkg/apis/redis/v1alpha1"
	customclient "github.com/wflysnow/my-middleware-operator/pkg/clients/clientset/versioned"
	custominformer "github.com/wflysnow/my-middleware-operator/pkg/clients/informers/externalversions/redis/v1alpha1"
	redisClusterLister "github.com/wflysnow/my-middleware-operator/pkg/clients/listers/redis/v1alpha1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinfomers "k8s.io/client-go/informers/apps/v1"
	clientset "k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/pkg/controller"
)

type RedisClusterOperator struct {
	customCRDClient customclient.Interface
	defaultClient   clientset.Interface
	kubeConfig      *rest.Config
	options         *options.OperatorManagerServer

	eventRecorder record.EventRecorder

	syncHandler func(dKey string) error

	enqueueRedisCluster func(redisCluster *redistype.RedisCluster)

	redisClusterInfomer cache.SharedIndexInformer

	redisClusterLister redisClusterLister.RedisClusterLister

	redisClusterListerSynced cache.InformerSynced

	stsLister appslisters.StatefulSetLister

	stsListerSynced cache.InformerSynced

	queue workqueue.RateLimitingInterface
}

func NewRedisClusterOperator(redisInformer custominformer.RedisClusterInformer, stsInformer appsinfomers.StatefulSetInformer, kubeClient clientset.Interface, customCRDClient customclient.Interface, kubeConfig *rest.Config, options options.OperatorManagerServer) (*RedisClusterOperator, error) {
	return nil, nil
}

func (rco *RedisClusterOperator) Run(workers int, stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer rco.queue.ShutDown()

	glog.Infof("starting rediscluster operator")
	defer glog.Infof("shutting down rediscluster operator")

	if !controller.WaitForCacheSync("rediscluster", stopCh, rco.redisClusterListerSynced, rco.stsListerSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(rco.worker, time.Second, stopCh)
	}

	<-stopCh

}

func (rco *RedisClusterOperator) worker() {
	for rco.processNextWorkItem() {

	}
}

func (rco *RedisClusterOperator) processNextWorkItem() bool {
	key, quit := rco.queue.Get()
	if quit {
		return false
	}
	// Done marks item as done process ing, and if it has been marked as dirty again
	// whild it was being processed, it will be re-added to teh queue for re-processing

	/*defer rco.queue.Done(key)
	err := rco.syncHandler(key.(string))
	rco.handlerErr(err,key)*/
	go rco.syncHandler(key.(string))
	return true
}
