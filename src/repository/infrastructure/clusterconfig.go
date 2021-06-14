package infrastructure

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var (
	clientset      kubernetes.Interface
	IClusterConfig ClusterConfigInterface = &ClusterConfig{}
)

type ClusterConfig struct{}

type ClusterConfigInterface interface {
	RetrieveClientSet() kubernetes.Interface
	ClusterConfig() kubernetes.Interface
}

func init() {
	clientset = IClusterConfig.ClusterConfig()
}

func ProvideClusterConfig() *ClusterConfig {
	return &ClusterConfig{}
}

func (clusterConfig *ClusterConfig) RetrieveClientSet() kubernetes.Interface {
	return clientset
}

func (clusterConfig *ClusterConfig) ClusterConfig() kubernetes.Interface {
	var clusterConfiguration *rest.Config
	if os.Getenv("IN_CLUSTER") == "true" {
		clusterConfiguration = retrieveInClusterConfig()
	} else {
		setKubeconfigFlag()
		clusterConfiguration = retrieveOutOfClusterConfig()
	}

	clientset, err := kubernetes.NewForConfig(clusterConfiguration)
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
}
