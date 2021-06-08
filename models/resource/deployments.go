package resource

type DeploymentStatus struct {
	Replicas            int32 `json:"replicas" xml:"replicas" form:"replicas" query:"replicas"`
	UpdatedReplicas     int32 `json:"updateRplicas" xml:"updateRplicas" form:"updateRplicas" query:"updateRplicas"`
	ReadyReplicas       int32 `json:"readyReplicas" xml:"readyReplicas" form:"readyReplicas" query:"readyReplicas"`
	AvailableReplicas   int32 `json:"availableReplicas" xml:"availableReplicas" form:"availableReplicas" query:"availableReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas" xml:"unavailableReplicas" form:"unavailableReplicas" query:"unavailableReplicas"`
}

type Deployment struct {
	Name      string           `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string           `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status    DeploymentStatus `json:"status" xml:"status" form:"status" query:"status"`
	Age       int64            `json:"age" xml:"age" form:"age" query:"age"`
}

type ResourceName string

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

type ResourceList map[ResourceName]string

type ResourceRequirements struct {
	Limits   ResourceList `json:"limits" xml:"limits" form:"limits" query:"limits"`
	Requests ResourceList `json:"requests" xml:"requests" form:"requests" query:"requests"`
}

type EnvVar struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Value string `json:"value" xml:"value" form:"value" query:"value"`
}

type VolumeMount struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	ReadOnly  bool   `json:"readOnly" xml:"readOnly" form:"readOnly" query:"readOnly"`
	MountPath string `json:"mountPath" xml:"mountPath" form:"mountPath" query:"mountPath"`
	SubPath   string `json:"subPath" xml:"subPath" form:"subPath" query:"subPath"`
}

type DeploymentParams struct {
	Name         string               `json:"name" xml:"name" form:"name" query:"name"`
	Image        string               `json:"image" xml:"image" form:"image" query:"image"`
	Resources    ResourceRequirements `json:"resources" xml:"resources" form:"resources" query:"resources"`
	Replicas     int32                `json:"replicas" xml:"replicas" form:"replicas" query:"replicas"`
	Env          []EnvVar             `json:"env" xml:"env" form:"env" query:"env"`
	VolumeMounts []VolumeMount        `json:"volumeMounts" xml:"volumeMounts" form:"volumeMounts" query:"volumeMounts"`
	NodeName     string               `json:"nodeName" xml:"nodeName" form:"nodeName" query:"nodeName"`
}