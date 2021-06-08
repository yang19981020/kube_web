package resource

import (
	"context"
	"encoding/json"
	"kube_web/models/resource"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ks K8SService) CreatePVC(namespace string, pvcs *resource.PVCParams) (*v1.PersistentVolumeClaim, error) {
	var resource v1.ResourceList

	// 组织资源数据
	resourceParams, _ := json.Marshal(pvcs.Capacity)
	err := json.Unmarshal(resourceParams, &resource)
	if err != nil {
		return nil, err
	}
	// 组织labels
	labels := map[string]string{"app": pvcs.Name}
	// 组织选择器
	var selector = metav1.LabelSelector{MatchLabels: labels}

	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvcs.Name,
			Labels: labels,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			Selector: &selector,
		},
	}

	kpvc, err := ks.clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	return kpvc, err
}
