package main

import (
	"fmt"
)

func main () {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	
	var comando int
	fmt.Scanf("%d", &comando)
}