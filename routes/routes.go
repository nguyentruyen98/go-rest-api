package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events", getEvents)
	authenticated.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
