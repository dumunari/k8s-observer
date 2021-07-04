package client

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

type Resources struct {
	Config kubernetes.Interface
}

type ResourcesInterface interface {
	RetrieveDeploymentList() (*v1.DeploymentList, error)
	RetrieveNodesList() (*v12.NodeList, error)
}

func (resourcesReceiver *Resources) RetrieveDeploymentList() (*v1.DeploymentList, error){
	log.Println("[Resources] - RetrieveDeploymentList")
	return resourcesReceiver.Config.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
}

func (resourcesReceiver *Resources) RetrieveNodesList() (*v12.NodeList, error) {
	log.Println("[Resources] - RetrieveNodesList")
	return resourcesReceiver.Config.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

