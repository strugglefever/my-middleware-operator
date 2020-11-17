package v1alpha1

import (
	"fmt"
	"testing"
)

func TestServerUp(t *testing.T) {
	redisCluster := &RedisCluster{}
	fmt.Printf("init redisCluster:%#v\n", &redisCluster)
	leaderElectionConfiguration := &LeaderElectionConfiguration{}
	fmt.Printf("init leaderElectionConfigration:%#v\n",leaderElectionConfiguration)
	operatorManagerConfig := &OperatorManagerConfig{}
	fmt.Printf("init operatorManagerConfig: %#v\n", operatorManagerConfig)
}
