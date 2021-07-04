package providers

import (
	"observer/src/api/repositories"
	"observer/src/api/services"
	"observer/src/api/utils"
	"observer/src/infrastructure/client"
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

func ProvideConfig() *client.Config{
	return &client.Config{}
}