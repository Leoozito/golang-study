package controller

import (
	"github.com/api/model"
	"github.com/gin-gonic/gin"
)

func GetAlunos(c *gin.Context) {
	c.JSON(200, model.Alunos)
}

func GetAlunoById(c *gin.Context) {
	id := c.Params.ByName("id")
	
	c.JSON(200, gin.H{
		"id": id,
	})
}