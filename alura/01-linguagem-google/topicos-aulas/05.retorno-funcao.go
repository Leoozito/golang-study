// Esse arquivo mostra: funcionamento de RECEBER e RETORNAR mais de uma coisa em uma função
// Mostra em todos detalhes como trabalhar com FUNÇÕES

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// recebendo mais + 1 dado no retorno 
	altura, peso := devolveAlturaePeso()
	fmt.Println("Tenho ", altura, " de altura e ", peso, " kilos\n")

	// ignorando um retorno
	nome, _ := devolveNomeeIdade()

	// recebendo um retorno
	comando := leComando()

	switch comando {
	case 1:
		// chama função e envia dados por parametro
		iniciarMonitoramento(nome);
	case 0:
		fmt.Println("Saindo do Programa !")
		os.Exit(0)
	default:
		fmt.Println("Comando invalido")
	}
}

// função que retorna + de 1 valor
func devolveAlturaePeso() (float64, int) {
	altura := 1.74
	peso := 20
	return altura, peso
}

// função que retorna + de 1 valor
func devolveNomeeIdade() (string, int) {
	nome := "Leonardo"
	idade := 20
	return nome, idade
}

// função que retorna um valor
// retorna valor inserido pelo úsuario
func leComando() int {
	fmt.Println("Digite o comando desejado")
	fmt.Println("1- Iniciar monitoramento de site")
	fmt.Println("0 - Sair do programa\n")

	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("\nO comando escolhido foi", comandoLido)

	return comandoLido
}

// recebe dados por parametro
func iniciarMonitoramento(nome string) {
	versao := 1.1
	fmt.Println("\nOlá, sr", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"

	// ignorando um retorno
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso\n")
	} else {
		fmt.Println("Site", site, "esta com problemas. StatusCode:", resp.StatusCode, "\n")
	}
}
