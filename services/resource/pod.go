package resource

import (
	"context"
	"encoding/json"
	"kube_web/models/resource"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ks K8SService) ListPod(namespace string) (*v1.PodList, error) {
	pods, err := ks.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return pods, err
}

func (ks K8SService) GetPod(namespace, pod string) (*v1.Pod, error) {
	kpod, err := ks.clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	return kpod, err
}

func (ks K8SService) DeletePod(namespace, pod string) error {
	err := ks.clientset.CoreV1().Pods(namespace).Delete(context.TODO(), pod, metav1.DeleteOptions{})
	return err
}

func (ks K8SService) CreatePod(namespace string, pps *resource.PodParams) (*v1.Pod, error) {
	var env []v1.EnvVar
	var resource v1.ResourceRequirements

	// 组织环境变量
	envParams, _ := json.Marshal(pps.Env)
	err := json.Unmarshal(envParams, &env)
	if err != nil {
		return nil, err
	}
	// 组织资源数据
	resourceParams, _ := json.Marshal(pps.Resources)
	err = json.Unmarshal(resourceParams, &resource)
	if err != nil {
		return nil, err
	}
	// 组织labels
	labels := map[string]string{"run": pps.Name}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pps.Name,
			Labels: labels,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            pps.Name,
					Image:           pps.Image,
					ImagePullPolicy: "Always",
					Env:             env,
					Resources:       resource,
				},
			},
			NodeName: pps.NodeName,
		},
	}

	kpod, err := ks.clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	return kpod, err
}
