package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MiddlewareRedisTypeKey        = "redis.middleware.hc.cn"
	MiddlewareRedisClustersPrefix = "redisclusters-"
)

type RedisClusterUpdateStrategyType string

const (
	AssginReceiveStrategyType RedisClusterUpdateStrategyType = "AssginReceive"
	AutoReceiveStrategyType   RedisClusterUpdateStrategyType = "AutoReceive"
)

type RedisClusterConditionType string

const (
	MasterConditionType RedisClusterConditionType = "master"
	SlaveConditionType  RedisClusterConditionType = "slave"
)

type RedisClusterPhase string

// These are the valid phases of a RedisCluster
const (
	RedisClusterUpgrading RedisClusterPhase = "Upgrading"
	// RedisClusterNone means the RedisCluster crd is first create
	RedisClusterNone RedisClusterPhase = "None"
	//RedisClusterCreatin means the RedisCluster is Creating
	RedisClusterCreating RedisClusterPhase = "Creating"
	RedisClusterRunning  RedisClusterPhase = "Running"
	RedisClusterFailed   RedisClusterPhase = "Failed"
	RedisClusterdScaling RedisClusterPhase = "Scaling"
	RedisClusterDeleting RedisClusterPhase = "Deleting"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RedisCluster struct {
	metav1.TypeMeta `json:",inline"`
	// TODO note should be
	// +optional
	metav1.ObjectMeta `json;"metadata,omitempty"`
	// +optional
	Spec RedisClusterSpec `json:"spec,omitempty"`
	// +optional
	Status RedisClusterStatus `json:"status,omitempty"`
}

// RedisClusterSpec describes the desired functionality of the complexPodScale
type RedisClusterSpec struct {
	Replicas *int32 `json:"replicas,omitempty"`

	Pause          bool                          `json:"pause,omitempty"`
	Finalizers     string                        `json:"finalizers,omitempty"`
	Repository     string                        `json:"repository,omitempty"`
	Version        string                        `json:"version,omitempty"`
	UpdateStrategy RedisClusterUpdateStrategy    `json:"updateStrategy,omitempty"`
	Pod            []RedisClusterPodTemplateSpec `json:"pod,omitempty"`
}

type RedisClusterUpdateStrategy struct {
	Type             RedisClusterUpdateStrategyType `json:"type,omitempty"`
	Pipeline         string                         `json:"pipeline,omitempty"`
	AssignStrategies []SlotsAssignStrategy          `json:"assignStrategies,omitempty"`
}

type SlotsAssignStrategy struct {
	Slots *int32 `json:"slots,omitempty"`
	//nodeid
	FromReplicas string `json:"fromReplicas,omitempty"`
}

type RedisClusterPodTemplateSpec struct {
	Configmap       string                `json:"configmap,omitempty" protobuf:"bytes,4,opt,name=configmap"`
	MonitorImage    string                `json:"monitorImage,omitempty" protobuf:"bytes,4,opt,name=monitorImage"`
	InitImage       string                `json:"initImage,omitempty" protobuf:"bytes,4,opt,name=initImage"`
	MiddlewareImage string                `json:"midddlewareImage,omitempty" protobuf:"bytes,4,opt,name=middlewareImage"`
	Volumes         RedisClusterPodVolume `json:"volumes"`

	// list of environment variables to set in the container
	// cannot be update
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Env []v1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
	// +optional
	Resources v1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,4,opt,name=resources"`
	// +optional
	Affinity       *v1.Affinity                     `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	Annotations    map[string]string                `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
	Labels         map[string]string                `json:"labels,omitempty" protobuf:"bytes,11,rep,name=,labels"`
	UpdateStrategy appsv1.StatefulSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,7,opt,name=updateStrategy"`
}

type RedisClusterPodVolume struct {
	Type                      string `json:"type,omitempty" protobuf:"bytes,4,opt,name=type"`
	PersistentVolumeClaimNmae string `json:"persistentVolumeClaimNmae,omitempty" protobuf:"bytes,4,opt,name=persistentVolumeClaimName"`
}

type RedisClusterStatus struct {
	Replicas int32 `json:"replicas" protobuf:"varint,2,opt,name=replicas"`
	// +optional
	Reason     string                  `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Phase      RedisClusterPhase       `json:"phase"`
	Conditions []RedisClusterCondition `json:"conditions"`
}

type RedisClusterCondition struct {
	Name         string                    `json:"name,omitempty" protobuf:"bytes,4,opt,name=name"`
	Type         RedisClusterConditionType `json:"type"`
	Instance     string                    `json:"instance,omitempty" protobuf:"bytes,4,opt,name=instance"`
	NodeId       string                    `json:"nodeId,omitempty" protobuf:"bytes,4,opt,name=nodeId"`
	MasterNodeId string                    `json:"masterNodeId,omitempty" protobuf:"bytes,4,opt,name=masterNodeId"`
	DomainName   string                    `json:"domainName,omitempty" protobuf:"bytes,4,opt,name=domainName"`
	Slots        string                    `json:"slots,omitempty" protobuf:"bytes,4,opt,name=slots"`
	Hostname     string                    `json:"hostname,omitempty" protobuf:"bytes,4,opt,name=hostname"`
	HostIp       string                    `json:"hostIp,omitempty" protobuf:"bytes,4,opt,name=hostIp"`
	Status       v1.ConditionStatus        `json:status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io./api/core/v1.ConditionStatus"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// +optional
	Resson string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// +K8S:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata"`

	Items []RedisCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OperatorManagerConfig struct {
	metav1.TypeMeta

	// Operators is the list of operators to enable or disable
	// '*' means "all enable by default operators"
	// 'foo' means "enable 'foo'"
	// '-foo' means "disable 'foo'"
	// first item for a particular name wins
	Operators []string

	ConcurrentRedisClusterSyncs int32

	// cluster create or upgrade timeout(min)
	ClusterTimeOut int32

	// How long to wait between starting controller managers
	ControllerStartInterval metav1.Duration

	ResysncPeriod int64

	// leaderElection defines the configuration of leader election client.
	LeaderElection LeaderElectionConfiguration

	// port is the port that the controller manager's http service runs on.
	Port int32

	// address is the IP address to serve on (set to 0.0.0.0 for all interface).
	Address string

	// enableProfiling enables profiling via web interface host:port/debug/pprof
	EnableProfiling bool
	//contentType is contentType of requests sent to apiserver
	ContentType string
	// kubeAPIQPS is the QPS to use whild talking with kubernetes apiserver
	KubeAPIQPS float32
	// kubeAPIBurst is the burst to use while talking with kubernetes apiserver.
	KubeAPIBurst int32
}

// LeaderElectionConfiguration defines the configuration of leader election
// clients for components that can run with leader election enabled.
type LeaderElectionConfiguration struct {
	// leaderElect enables a leader election client to gain leadership
	// before executing the main loop. Enable this when running replicated
	// components for high availability.
	LeaderElect bool
	// leaseDuration is the duration that non-leader candidates will wait
	// after observing a leadership renewal until attempting to acquire
	// leadership of a led but unrenewed leader slot. This is effectively the
	// maximum duration that a leader can be stopped before it is replaced
	// by another candidate. This is only applicable if leader election is
	// enabled.
	LeaseDuration metav1.Duration
	// renewDeadline is the interval between attempts by the acting master to
	// renew a leadership slot before it stops leading. This must be less
	// than or equal to the lease duration. This is only applicable if leader
	// election is enabled.
	RenewDeadline metav1.Duration
	// retryPeriod is the duration the clients should wait between attempting
	// acquisition and renewal of a leadership. This is only applicable if
	// leader election is enabled.
	RetryPeriod metav1.Duration
	// resourceLock indicates the resource object type that will be used to lock
	// during leader election cycles.
	ResourceLock string
}
