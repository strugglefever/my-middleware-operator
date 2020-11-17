package v1alpha1

import (
	"fmt"
	"testing"
)

func TestServerUp(t *testing.T) {
	redisCluster := &RedisCluster{}
	fmt.Printf("init redisCluster:%#v", &redisCluster)
}
