package views

import "github.com/gin-gonic/gin"

var route *gin.Engine

func init () {
	route = gin.Default()
	BlogRouteGroup(route.Group("/api/blog"))
}

func GetRoute() *gin.Engine {
	return route
	
}
