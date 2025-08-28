package controller

import (
	"net/http"

	"github.com/api/database"
	"github.com/api/model"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	
	c.JSON(200, gin.H{
		"id": id,
	})
}

func CreateStudent(c *gin.Context) {
	var student model.Student
	// empacotar o corpo da requisição
	if err := c.ShouldBindJSON(&student)
	
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}