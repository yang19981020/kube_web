package resource

import (
	"context"
	"encoding/json"
	"kube_web/models/resource"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ks K8SService) ListDeployment(namespace string) (*appsv1.DeploymentList, error) {
	deployments, err := ks.clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	return deployments, err
}

func (ks K8SService) GetDeployment(namespace, deployment string) (*appsv1.Deployment, error) {
	kdeployment, err := ks.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment, metav1.GetOptions{})
	return kdeployment, err
}

func (ks K8SService) DeleteDeployment(namespace, deployment string) error {
	err := ks.clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment, metav1.DeleteOptions{})
	return err
}

func (ks K8SService) ApplyDeployment(namespace string, dps *resource.DeploymentParams, update bool) (*appsv1.Deployment, error) {
	var env []v1.EnvVar
	var resource v1.ResourceRequirements

	// 组织环境变量
	envParams, _ := json.Marshal(dps.Env)
	err := json.Unmarshal(envParams, &env)
	if err != nil {
		return nil, err
	}
	// 组织资源数据
	resourceParams, _ := json.Marshal(dps.Resources)
	err = json.Unmarshal(resourceParams, &resource)
	if err != nil {
		return nil, err
	}
	// 组织labels
	labels := map[string]string{"app": dps.Name}
	// 组织选择器
	var selector = metav1.LabelSelector{MatchLabels: labels}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dps.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &selector,
			Replicas: &dps.Replicas,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            dps.Name,
							Image:           dps.Image,
							ImagePullPolicy: "Always",
							Env:             env,
							Resources:       resource,
						},
					},
					NodeName: dps.NodeName,
				},
			},
		},
	}

	var kdeployment *appsv1.Deployment
	if update {
		kdeployment, err = ks.clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	} else {
		kdeployment, err = ks.clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	}

	return kdeployment, err
}
