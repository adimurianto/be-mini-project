package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.UserController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("user/", middleware.AuthMiddleware(), ctrl.GetData)
	groupRoutes.POST("user/", middleware.AuthMiddleware(), ctrl.CreateData)
}
