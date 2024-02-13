package routes

import (
	"github.com/gin-gonic/gin"
	"mario-mtz.com/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
