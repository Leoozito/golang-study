package main

import (
	"fmt"
	"net/http" // pacote do go responsavel por fazer acesso web
	"os"
)

func main() {

	for {
		menu()

		comando := get_comando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs: ")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando: [", comando, "] não identificado, verifique os comandos disponiveis nas opções")
		}
	}
}

func menu() {
	fmt.Println("----------------------------------")
	fmt.Println("Escolha alguma das opções abaixo")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func get_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando: [", comando, "] escolhido...")
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for _, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site: ", site, " em funcionamento, status: ", resp.StatusCode)
			fmt.Println("----------------------------------")
		} else {
			fmt.Println("Site: ", site, "status: ", resp.StatusCode, "ERROR")
			fmt.Println("----------------------------------")
		}
	}

}
