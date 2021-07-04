package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"path/filepath"
)

var (
	clientset kubernetes.Interface
	IConfig   ConfigInterface = &Config{}
)

type Config struct{}

type ConfigInterface interface {
	RetrieveClientSet() kubernetes.Interface
	ClientConfig() kubernetes.Interface
}

func init() {
	clientset = IConfig.ClientConfig()
}

func (clientConfigReceiver *Config) RetrieveClientSet() kubernetes.Interface {
	log.Println("[Config] - RetrieveClientSet")
	return clientset
}

func (clientConfigReceiver *Config) ClientConfig() kubernetes.Interface {
	log.Println("[Config] - ClientConfig")
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
	log.Println("[Config] - retrieveInClusterConfig")
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func retrieveOutOfClusterConfig() *rest.Config {
	log.Println("[Config] - retrieveOutOfClusterConfig")
	config, err := clientcmd.BuildConfigFromFlags("", flag.Lookup("kubeconfig").Value.(flag.Getter).Get().(string))
	if err != nil {
		panic(err.Error())
	}
	return config
}

func setKubeconfigFlag() {
	log.Println("[Config] - setKubeconfigFlag")
	home := homedir.HomeDir()
	flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
}
