package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/helm-service/pkg/models"
)

func GetHelmReleaseInfo(context *gin.Context) {
	releaseName := context.Param("releaseName")
	namespace := context.Param("namespace")
	// todo : read kubeconfig from header as encoded
	kubeConfigPath, err := models.GetKubeconfigPath(&context.Request.Header)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	releaseInfo, err := models.GetHelmReleaseInfo(releaseName, namespace, kubeConfigPath)
	models.RemoveTempFile(kubeConfigPath)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, releaseInfo)
}
