package main

import (
	"log"
	"net/http"
	"observer/src/router"
)

func main() {
	log.Println("k8s Observer Started")

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", r))

}
