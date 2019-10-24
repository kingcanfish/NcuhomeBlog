package views

import "github.com/gin-gonic/gin"

var route *gin.Engine

func init () {
	route = gin.Default()
	BlogGroup(route.Group("/api/blog"))
}
