package main

import (
	"log"
	"net/http"
	"observer/src/alerts"
	"observer/src/api/router"
)

func main() {
	log.Println("k8s Observer Started")

	alerts.StartAlertsSchedule()

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
