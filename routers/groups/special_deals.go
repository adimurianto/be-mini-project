package routers

import (
	controllers "be-mini-project/controllers"
	"be-mini-project/routers/middleware"

	"github.com/gin-gonic/gin"
)

func SpecialDealsRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.SpecialDealsController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("special-deals/", ctrl.GetData)
	groupRoutes.POST("special-deals/", middleware.AuthMiddleware(), ctrl.CreateData)
}
