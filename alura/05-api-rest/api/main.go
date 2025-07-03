package main

import (
	"github.com/api/model"
	"github.com/api/routes"
)

func main() {

	model.Alunos = []model.Aluno{
		{	
			Id: "1",
			Name: "Leonardo",
		},
		{	
			Id: "2",
			Name: "Lais",
		},
	}

	routes.Initialize()
}