package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/helm-service/pkg/models"
	"helm.sh/helm/v3/pkg/release"
)

var Info release.Info

func GetHelmReleaseInfo(context *gin.Context) {
	releaseName := context.Param("releaseName")
	namespace := context.Param("namespace")
	// todo : read kubeconfig from header as encoded
	kubeConfigPath := ""
	releaseInfo, err := models.GetHelmReleaseInfo(releaseName, namespace, kubeConfigPath)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	context.IndentedJSON(http.StatusOK, releaseInfo)
}
