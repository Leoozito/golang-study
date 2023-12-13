package main

import "fmt"

func main() {
	menu()

	comando := get_comando()

	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs: ")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Comando: [", comando, "] não identificado, verifique os comandos disponiveis nas opções")
	}
}

func menu() {
	fmt.Println("Escolha alguma das opções abaixo")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

// ao lado do parenteses da função: eu coloco o tipo do retorno que irei receber
func get_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando: [", comando, "] escolhido, verificando...")
	return comando
}
