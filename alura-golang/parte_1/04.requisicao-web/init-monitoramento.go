package main

import (
	"fmt"
	"os"
	"net/http"
)

func main () {

	exibeIntroducao()
	exibeMenu()

	_, idade := devolveNomeeIdade()
	fmt.Println("Tenho", idade, "anos")

	comando := leComando()

	switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa !")
			os.Exit(0)
		default:
			fmt.Println("Comando invalido")
	}
}

func devolveNomeeIdade() (string, int) {
	nome := "Leonardo"
	idade := 20
	return nome, idade
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

// a função vai receber apenas valor type int
func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido //para retornar para fora da função aquilo que eu consegui ler
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}