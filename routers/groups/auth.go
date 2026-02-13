package routers

import (
	controllers "be-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.AuthController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.POST("auth/login/", ctrl.Login)
}
