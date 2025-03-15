package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http" // pacote do go responsavel por fazer requisições web - GET, PSOT
	"os"
	"strings"
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
	sites := read_file()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Iniciando monitoramento...")
		for _, site := range sites {
			tests_sites(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func tests_sites(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " em funcionamento, status: ", resp.StatusCode)
		fmt.Println("")
	} else {
		fmt.Println("Site: ", site, ", status: ", resp.StatusCode, "ERROR")
		fmt.Println("")
	}
}

// lê todos os sites contidos dentro do arquivo
func read_file() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") // abrimos o arquivo para ser lido
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for { // loop para poder ler o arquivo o inteiro [...]

		dados, err := leitor.ReadString('\n') // [...] batendo em todas as quebras de linhas que o arquivo tiver

		if err == io.EOF { // quando terminar de ler o arquivo inteiro
			break // para a repetição
		}

		dados = strings.TrimSpace(dados) // tirando espaços, para não serem adicionado na lista
		sites = append(sites, dados)
	}

	arquivo.Close()

	return sites
}
