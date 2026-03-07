package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.OrderController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("order/", ctrl.GetData)
	groupRoutes.POST("order/", ctrl.CreateData)
	groupRoutes.PUT("order/", middleware.AuthMiddleware(), ctrl.UpdateData)
	groupRoutes.DELETE("order/:id", middleware.AuthMiddleware(), ctrl.DeleteData)
}
