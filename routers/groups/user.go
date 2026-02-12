package routers

import (
	controllers "be-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.UserController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("user/", ctrl.GetData)
	groupRoutes.POST("user/", ctrl.CreateData)
}
