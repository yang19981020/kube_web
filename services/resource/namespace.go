package resource

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ks K8SService) ListNamespace() (*v1.NamespaceList, error) {
	namespaces, err := ks.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	return namespaces, err
}

func (ks K8SService) GetNamespace(name string) (*v1.Namespace, error) {
	namespace, err := ks.clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	return namespace, err
}

func (ks K8SService) CreateNamespace(name string) (*v1.Namespace, error) {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	knamespace, err := ks.clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return knamespace, err
}

func (ks K8SService) DeleteNamespace(name string) error {
	err := ks.clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}