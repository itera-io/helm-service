package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/helm-service/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterHelmReleaseRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
