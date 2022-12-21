package main

import (
	"log"
	"net/http"

	"github.com/dumunari/k8s-observer/src/alerts"
	"github.com/dumunari/k8s-observer/src/api/router"
)

func main() {
	log.Println("k8s Observer Started!")

	alerts.StartAlertsSchedule()

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
