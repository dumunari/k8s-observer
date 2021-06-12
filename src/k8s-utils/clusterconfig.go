package k8s_utils

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var clientset *kubernetes.Clientset

func init() {
	clientset = ClusterConfig()
}

func RetrieveClientSet() *kubernetes.Clientset {
	return clientset
}

func ClusterConfig() *kubernetes.Clientset {
	var clusterConfig *rest.Config

	if os.Getenv("IN_CLUSTER") == "true" {
		clusterConfig = retrieveInClusterConfig()
	} else {
		setKubeconfigFlag()
		clusterConfig = retrieveOutOfClusterConfig()
	}

	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func retrieveInClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func retrieveOutOfClusterConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", flag.Lookup("kubeconfig").Value.(flag.Getter).Get().(string))
	if err != nil {
		panic(err.Error())
	}
	return config
}

func setKubeconfigFlag() {
	home := homedir.HomeDir()
	flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	flag.Parse()
}
