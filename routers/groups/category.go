package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.CategoryController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("category/", ctrl.GetData)
	groupRoutes.POST("category/", middleware.AuthMiddleware(), ctrl.CreateData)
	groupRoutes.PUT("category/", middleware.AuthMiddleware(), ctrl.UpdateData)
	groupRoutes.DELETE("category/:id", middleware.AuthMiddleware(), ctrl.DeleteData)
}
