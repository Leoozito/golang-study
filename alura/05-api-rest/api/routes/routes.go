package routes

import (
	"github.com/api/controller"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/students", controller.GetStudents)
	r.GET("/student/:id", controller.GetStudentById)
	r.POST("/student", controller.CreateStudent)
	r.Run(":8000")
}