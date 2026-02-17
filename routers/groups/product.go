package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.ProductController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("product/", ctrl.GetData)
	groupRoutes.POST("product/", middleware.AuthMiddleware(), ctrl.CreateData)
}
