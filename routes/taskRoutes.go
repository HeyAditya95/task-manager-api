// Author: https://github.com/HeyAditya95

package routes

import (
	"github.com/HeyAditya95/task-manager-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine) {
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
