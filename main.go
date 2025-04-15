// Author: https://github.com/HeyAditya95

package main

import (
	"fmt"

	"github.com/HeyAditya95/task-manager-api/database"
	"github.com/HeyAditya95/task-manager-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Starting Task Manager API...")

	database.ConnectDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.RegisterTaskRoutes(r)

	r.Run(":8080")
}
