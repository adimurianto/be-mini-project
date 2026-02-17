package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func SpecialDealsItemRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.SpecialDealsItemController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("special-deals-items/", ctrl.GetData)
	groupRoutes.POST("special-deals-items/", middleware.AuthMiddleware(), ctrl.CreateData)
}
