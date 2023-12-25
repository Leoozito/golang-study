package main

import (
	"fmt"
	"net/http" // pacote do go responsavel por fazer acesso web
	"os"
	"time"
)

const delay = 5
const monitoramentos = 3

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
	fmt.Println("Escolha alguma das opções abaixo")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("")
}

func get_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando: [", comando, "] escolhido...")
	fmt.Println("")
	return comando
}

func iniciarMonitoramento() {
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Iniciando monitoramento...")
		for _, site := range sites {
			tests_sites(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func tests_sites(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " em funcionamento, status: ", resp.StatusCode)
		fmt.Println("")
	} else {
		fmt.Println("Site: ", site, "status: ", resp.StatusCode, "ERROR")
		fmt.Println("")
	}
}
