package routes

import (
	"github.com/api/controller"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/alunos", controller.GetAlunos)
	r.GET("/alunos/:id", controller.GetAlunoById)
	r.Run("8000")
}