package services_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"observer/src/services"
	"os"
	"testing"
)

var (
	getDeployments func() (*v1.DeploymentList, error)
)

type getRepositoryMock struct {
	mock.Mock
}

func (m *getRepositoryMock) RetrieveDeployments() (*v1.DeploymentList, error) {
	m.Called()
	return getDeployments()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestRetrieveServicesService_2_deployments_found(t *testing.T) {
	//arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "alphaTest",
							"service": "service-test-1",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "alphaTest",
									"service": "service-test-1",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 2},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "alphaTest",
							"service": "service-test-2",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "alphaTest",
									"service": "service-test-2",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 2},
				},
			},
		}, nil
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesService()

	//assert
	assert.NotNil(t, services)
	assert.EqualValues(t, 2, len(services))
	assert.Nil(t, err)
}

func TestRetrieveServicesService_error(t *testing.T) {
	// arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{},
		}, errors.New("repository error")
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesService()

	//assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "repository error", fmt.Sprint(err))
	assert.Nil(t, services)
}

func TestRetrieveServicesService_0_deployments_found(t *testing.T) {
	//assert
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{},
		}, nil
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesService()

	//assert
	assert.Nil(t, services)
	assert.EqualValues(t, 0, len(services))
	assert.Nil(t, err)
}

func TestRetrieveServicesByApplicationGroupService_2_deployments_found_1_filter_application_group_beta(t *testing.T) {
	//arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "alphaTest",
							"service": "service-test-1",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "alphaTest",
									"service": "service-test-1",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 2},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "betaTest",
							"service": "service-test-2",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "betaTest",
									"service": "service-test-2",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 3},
				},
			},
		}, nil
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesByApplicationGroupService("betaTest")

	//assert
	assert.NotNil(t, services)
	assert.EqualValues(t, 1, len(services))
	assert.Nil(t, err)
}

func TestRetrieveServicesByApplicationGroupService_2_deployments_found_1_filter_no_application_group(t *testing.T) {
	//arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"service": "service-test-1",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"service": "service-test-1",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 2},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "betaTest",
							"service": "service-test-2",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "betaTest",
									"service": "service-test-2",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 3},
				},
			},
		}, nil
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesByApplicationGroupService("none")

	//assert
	assert.NotNil(t, services)
	assert.EqualValues(t, 1, len(services))
	assert.Nil(t, err)
}

func TestRetrieveServicesByApplicationGroupService_2_deployments_found_0_filter_non_existent_application_group(t *testing.T) {
	//arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"service": "service-test-1",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"service": "service-test-1",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 2},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"applicationGroup": "betaTest",
							"service": "service-test-2",
						},
					},
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"applicationGroup": "betaTest",
									"service": "service-test-2",
								},
							},
						},
					},
					Status: v1.DeploymentStatus{AvailableReplicas: 3},
				},
			},
		}, nil
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesByApplicationGroupService("delta")

	//assert
	assert.Nil(t, services)
	assert.EqualValues(t, 0, len(services))
	assert.Nil(t, err)
}

func  TestRetrieveServicesByApplicationGroup_error(t *testing.T) {
	// arrange
	getDeployments = func() (*v1.DeploymentList, error){
		return &v1.DeploymentList{
			Items: []v1.Deployment{},
		}, errors.New("repository error")
	}

	repository := &getRepositoryMock{}
	repository.On("RetrieveDeployments")

	service := &services.ServicesService{
		Repository: repository,
	}

	//act
	services, err := service.RetrieveServicesByApplicationGroupService("alpha")

	//assert
	assert.Nil(t, services)
	assert.EqualValues(t, 0, len(services))
	assert.NotNil(t, err)
	assert.NotNil(t, "repository error", fmt.Sprintf("%s", err))
}

func TestProvideServicesService(t *testing.T) {
	repository := &getRepositoryMock{}
	repository.On("ProvideDeploymentsRepository")

	result := services.ProvideServicesService()

	assert.NotNil(t, result)
	assert.IsType(t, services.ServicesService{}, *result)
}
