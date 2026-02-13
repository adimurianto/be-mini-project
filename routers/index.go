package routers

import (
	"net/http"

	routersGroup "be-mini-project/routers/groups"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	// Swagger Docs
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiVersion := "/api/v1/"
	routersGroup.AuthRoutes(route, apiVersion)
	routersGroup.UserRoutes(route, apiVersion)
	//Add All route
}
