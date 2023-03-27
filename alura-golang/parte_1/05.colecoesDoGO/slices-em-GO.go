package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

const monitoramento = 3
const delay = 2

func main () {
	iniciarMonitoramento()
	// exibeIntroducao()
	// exibeMenu()
	exibeNomes()

	// _, idade := devolveNomeeIdade()
	// fmt.Println("Tenho", idade, "anos")

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
	sites := []string {"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br" }
	
	// fmt.Println(sites)
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
	
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	nomes = append(nomes, "Aparecida") // o append ADICIONA dentro da SLICE
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes), "posições.") //ver quantos elementos tem
	fmt.Println("O meu slice tem", cap(nomes), "de capacidade.") // quanto de capacidade
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "esta com problemas. StatusCode:", resp.StatusCode)
	}

}