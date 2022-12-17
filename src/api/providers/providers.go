package providers

import (
	"github.com/dumunari/k8s-observer/src/api/repositories"
	"github.com/dumunari/k8s-observer/src/api/services"
	"github.com/dumunari/k8s-observer/src/api/utils"
	"github.com/dumunari/k8s-observer/src/infrastructure/client"
)

func ProvideResponseUtils() *utils.Response {
	return &utils.Response{}
}

func ProvideDeploymentsService() *services.DeploymentsService {
	return &services.DeploymentsService{
		Repository: ProvideDeploymentsRepository(),
	}
}

func ProvideNodesService() *services.NodesService {
	return &services.NodesService{
		Repository: ProvideNodesRepository(),
	}
}

func ProvideDeploymentsRepository() *repositories.DeploymentsRepository {
	return &repositories.DeploymentsRepository{
		Resources: ProvideResource(),
	}
}

func ProvideNodesRepository() *repositories.NodesRepository {
	return &repositories.NodesRepository{
		Resources: ProvideResource(),
	}
}

func ProvideResource() *client.Resources {
	return &client.Resources{
		Config: ProvideConfig().RetrieveClientSet(),
	}
}

func ProvideConfig() *client.Config {
	return &client.Config{}
}
