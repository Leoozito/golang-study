// Fiz nesse arquivo uma aplicação para rodar. Nessa aplicação mostra perfeitamente como slice armazena os valores sem precisar definir quantia máxima de dados.

package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		working_with_slices()
	}
}

func working_with_slices() []string {
	slice := []string{}
	var data string

	apresentation()

	fmt.Scanln(&data)
	slice = append(slice, data)

	fmt.Println("Dados do slice: ", slice)
	fmt.Println("Quantidade armazenada: ", len(slice))
	fmt.Println("----------------------------------------------------")

	question_continue()

	return slice
}

func apresentation() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("Digite algo para ser adicionado dentro da slice")
	fmt.Println("----------------------------------------------------")
	fmt.Println(" ")
}

func question_continue() {
	fmt.Println(" ")
	fmt.Println("Quer continuar?")
	fmt.Println("1 - Sim")
	fmt.Println("0 - Não")
	fmt.Println(" ")

	var resp int
	fmt.Scan(&resp)
	switch resp {
	case 0:
		os.Exit(0)
	}
}
