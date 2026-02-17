package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func BannerRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.BannerController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("banner/", ctrl.GetData)
	groupRoutes.POST("banner/", middleware.AuthMiddleware(), ctrl.CreateData)
}
