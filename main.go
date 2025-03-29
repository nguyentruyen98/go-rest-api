package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/db"
	"github.com/lolstate/lol-api-service/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
