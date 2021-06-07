package utils

import (
	"flag"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)
var (
	config *string
	kub_config_path string
	err error
	kube_config *rest.Config
	client *kubernetes.Clientset
)
func newClient() *kubernetes.Clientset {
	if kub_config_path,err = beego.AppConfig.String("kubeconfigpath"); err != nil {
		panic(err.Error())
	}
	config = flag.String("kubeconfig", kub_config_path, "absolute path to the kubeconfig file")
	flag.Parse()
	kube_config, err = clientcmd.BuildConfigFromFlags("", *config)
	if err != nil {
		panic(err.Error())
	}
	client, err = kubernetes.NewForConfig(kube_config)
	if err != nil {
		panic(err.Error())
	}
	return client
}
func getRestConf()  (kube_conf *rest.Config) {
	var (
		kubeconfig []byte
		err error
	)
	if kub_config_path,err = beego.AppConfig.String("kubeconfigpath"); err != nil {
		panic(err.Error())
	}
	if kubeconfig, err = ioutil.ReadFile(kub_config_path); err != nil {
		panic(err.Error())
	}
	if kube_conf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
		panic(err.Error())
	}
	return kube_conf
}
var K8sClient = newClient()
var K8sConfig = getRestConf()