package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func LogoRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.LogoController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("logo/", ctrl.GetData)
	groupRoutes.POST("logo/", middleware.AuthMiddleware(), ctrl.CreateData)
}
