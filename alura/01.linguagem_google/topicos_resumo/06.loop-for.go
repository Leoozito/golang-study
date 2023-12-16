package main

import (
	"fmt"
	"os"
)

func main() {
	//  no GO não tem "WHILE", ou seja, usando o for pode se fazer loop infinito
	for {
		exibeMenu()

		comando := retornaComando()

		switch comando {
		case 1:
			FOR_with_RANGE()
		case 2:
			fmt.Println("Monitorando...")
		case 3:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa !")
			os.Exit(0)
		default:
			fmt.Println("Comando invalido")
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Ver funcionamento do FOR com RANGE [adicional importante]")
	fmt.Println("2- Iniciar Monitoramento")
	fmt.Println("3- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func retornaComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func FOR_with_RANGE() {
	nomes := []string{"Leonardo", "Lais", "Vitoria"}
	fmt.Println("--------------------------------------------------")
	fmt.Println(" LOOP FOR COM RANGE ")
	fmt.Println(" ")
	fmt.Println("Essa é sua lista: ", nomes)

	for i, nome := range nomes { // o range do for é IMPORTANTE:
		// com o range, é possivel acessar as posições da lista (no caso é o "i") e o valor das posições
		fmt.Println("O nome: ", nome, "está na posição: ", i)
	}

	fmt.Println(" ")
	fmt.Println("--------------------------------------------------")
}
