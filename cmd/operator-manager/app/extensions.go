package app

import (
	"fmt"

	"github.com/wflysnow/my-middleware-operator/pkg/operator/redis"
)

func startRedisClusterController(otx OperatorContext) (bool, error) {
	rco, err := redis.NewRedisClusterOperator(
		otx.RedisInformerFactory.Cr().V1alpha1().RedisClusters(),
		otx.InformerFactory.Apps().V1().StatefulSets(),
		otx.DefaultClientBuilder.ClientOrDie("default-kube-client"),
		otx.OperatorClientBuilder.ClientOrDie("redisculster-operator"),
		otx.kubeConfig,
		otx.Options,
	)
	if err != nil {
		return true, fmt.Errorf("error creating rediscluster operator:%v", err)
	}
	go rco.Run(int(otx.Options.ConcurrentRedisClusterSyncs), otx.Stop)
	return true, nil
}
