package resource

import (
	"context"
	"encoding/json"
	"kube_web/models/resource"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ks K8SService) ListService(namespace string) (*v1.ServiceList, error) {
	services, err := ks.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	return services, err
}

func (ks K8SService) GetService(namespace, name string) (*v1.Service, error) {
	service, err := ks.clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	return service, err
}

func (ks K8SService) CreateService(namespace string, sp *resource.ServiceParams) (*v1.Service, error) {
	var svct v1.ServiceType

	// 组织服务类型
	typeParams, _ := json.Marshal(sp.Type)
	err := json.Unmarshal(typeParams, &svct)
	if err != nil {
		return nil, err
	}
	// 组织Port

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: sp.Name,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{"app": sp.TargetName},
			Type:     svct,
			Ports: []v1.ServicePort{
				{
					Name:     "http",
					Port:     sp.Port,
					Protocol: "TCP",
				},
			},
		},
	}

	kservice, err := ks.clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	return kservice, err
}

func (ks K8SService) DeleteService(namespace, service string) error {
	err := ks.clientset.CoreV1().Services(namespace).Delete(context.TODO(), service, metav1.DeleteOptions{})
	return err
}
