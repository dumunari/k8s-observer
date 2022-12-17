package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/dumunari/k8s-observer/src/api/services"
	"github.com/stretchr/testify/assert"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestRetrieveModulesService_2_deployments_found(t *testing.T) {
	//arrange
	retrieveDeployments = func() (*v12.DeploymentList, error) {
		return &v12.DeploymentList{
			Items: []v12.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"service": "service-test-1",
						},
					},
					Spec: v12.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"service": "service-test-1",
								},
							},
						},
					},
					Status: v12.DeploymentStatus{AvailableReplicas: 2},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"service": "service-test-2",
						},
					},
					Spec: v12.DeploymentSpec{
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"service": "service-test-2",
								},
							},
						},
					},
					Status: v12.DeploymentStatus{AvailableReplicas: 2},
				},
			},
		}, nil
	}

	repository := &deploymentsRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.DeploymentsService{
		Repository: repository,
	}

	//act
	modules, err := service.RetrieveDeployments()

	//assert
	assert.NotNil(t, modules)
	assert.EqualValues(t, 2, len(modules))
	assert.Nil(t, err)
}

func TestRetrieveModulesService_error(t *testing.T) {
	// arrange
	retrieveDeployments = func() (*v12.DeploymentList, error) {
		return &v12.DeploymentList{
			Items: []v12.Deployment{},
		}, errors.New("repository error")
	}

	repository := &deploymentsRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.DeploymentsService{
		Repository: repository,
	}

	//act
	modules, err := service.RetrieveDeployments()

	//assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "repository error", fmt.Sprint(err))
	assert.Nil(t, modules)
}

func TestRetrieveModulesService_0_deployments_found(t *testing.T) {
	//assert
	retrieveDeployments = func() (*v12.DeploymentList, error) {
		return &v12.DeploymentList{
			Items: []v12.Deployment{},
		}, nil
	}

	repository := &deploymentsRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.DeploymentsService{
		Repository: repository,
	}

	//act
	modules, err := service.RetrieveDeployments()

	//assert
	assert.Nil(t, modules)
	assert.EqualValues(t, 0, len(modules))
	assert.Nil(t, err)
}
