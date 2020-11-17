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

type LeaderElectionConfiguration struct {
	LeaderElect bool
	LeaseDuration metav1.Duration
	RenewDeadline metav1.Duration
	RetryPeriod metav1.Duration
	ResourceLock string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OperatorManagerConfig struct {
	metav1.TypeMeta
	Operators []string
	ConcurrentRedisClusterSyncs int32
	ClusterTimeOut int32
	ControllerStartInterval metav1.Duration
	ResysncPeriod int64
	LeaderElection LeaderElectionConfiguration
	Port int32
	Address string
	EnableProfiling bool
	ContentType string
	KubeAPIQPS float32
	KubeAPIBurst int32
}

