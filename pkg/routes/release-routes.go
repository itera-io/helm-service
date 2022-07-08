package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itera-io/helm-service/pkg/controllers"
)

var RegisterHelmReleaseRoutes = func(router *gin.Engine) {
	router.GET("/release-info/:releaseName/:namespace", controllers.GetHelmReleaseInfo)
}
