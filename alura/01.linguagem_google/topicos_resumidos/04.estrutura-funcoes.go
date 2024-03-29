// Este arquivo mostra: Organização de funções separadamente

package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

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
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

// ao lado do parenteses da função: eu coloco o tipo do retorno que irei receber
func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
