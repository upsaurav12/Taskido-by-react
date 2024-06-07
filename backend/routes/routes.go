package routes

import (
	"taskido/backend/controllers"

	"github.com/gin-gonic/gin"
)

func Setuproutes() *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome bruh!!!")
	})

	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	return r
}
