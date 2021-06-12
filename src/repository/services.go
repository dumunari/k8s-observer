package repository

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	k8s_utils "observer/src/k8s-utils"
)

func RetrieveDeploymentsRepository() (*v1.DeploymentList, error) {
	log.Println("RetrieveDeploymentsRepository")

	deployments, err := k8s_utils.RetrieveClientSet().AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployments %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, err
	}

	return deployments, nil
}
