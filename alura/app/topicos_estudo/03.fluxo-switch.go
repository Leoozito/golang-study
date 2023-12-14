package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa !")
	// } else {
	// 	fmt.Println("Comando invalido")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa !")
		os.Exit(0)
	default:
		fmt.Println("Comando invalido")
	}

	// No GO não precisa colocar o break
}
