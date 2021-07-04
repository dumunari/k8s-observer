package alerts

import (
	"github.com/robfig/cron"
	"log"
	"observer/src/api/providers"
	"observer/src/infrastructure/client"
)

var resources client.Resources

func init() {
	resources = client.Resources{
		Config: providers.ProvideConfig().RetrieveClientSet(),
	}
}

func StartAlertsSchedule(){
	log.Println("[Alerts] - StartAlertsSchedule")

	c := cron.New()
	err := c.AddFunc("@every 5s", func() {
		DeploymentsCheck()
		NodesCheck()
	})
	if err != nil {
		log.Println("[Alerts] - Alerts couldn't be scheduled.")
		return
	}

	c.Start()
}
