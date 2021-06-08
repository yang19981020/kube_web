package resource

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kube_web/utils"
)

type K8SService struct {
	clientset *kubernetes.Clientset
}

func New() K8SService {
	config := utils.K8sConfig
	// create the clientset: *kubernetes.Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ks := K8SService{clientset}
	return ks
}
func (ks K8SService) ListNode() {
	// get nodes
	nodes, err := ks.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	for i, node := range nodes.Items {
		fmt.Println(i, node.Name)
	}
}
