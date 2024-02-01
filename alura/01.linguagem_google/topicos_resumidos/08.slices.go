// Fiz nesse arquivo uma aplicação para rodar. Nessa aplicação mostra perfeitamente como slice armazena os valores sem precisar definir quantia máxima de dados.

package main

import (
	"fmt"
	"os"
)

func main() {
	slice := []string{}

	for {
		slice = working_with_slices(slice)
	}
}

func working_with_slices(slice []string) []string {
	var data string

	apresentation()

	fmt.Scanln(&data)
	slice = append(slice, data)
	fmt.Println(" ")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Dados do slice: ", slice)
	fmt.Println("Quantidade armazenada: ", len(slice))

	question_continue(slice)

	return slice
}

func apresentation() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("Digite algo para ser adicionado dentro da slice")
	fmt.Println("----------------------------------------------------")
	fmt.Println(" ")
}

func question_continue(slice []string) {
	fmt.Println(" ")
	fmt.Println("Quer continuar?")
	fmt.Println("1 - Sim")
	fmt.Println("0 - Não")
	fmt.Println(" ")

	var resp string
	fmt.Scan(&resp)
	switch resp {
	case "0":
		fmt.Println("Saindo ...")
		os.Exit(0)
	case "1":
		working_with_slices(slice)
	default:
		fmt.Println("Comando indefinido, escolha entre as opções apresentadas")
		question_continue(slice)
	}

}
