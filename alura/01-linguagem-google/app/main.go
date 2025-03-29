package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http" // pacote do go responsavel por fazer requisições web - GET, PSOT
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5
const monitoramentos = 3

func main() {
	for {
		menu()
		comando := getComand()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs: ")
			showLogs()
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

func getComand() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando: [", comando, "] escolhido...")
	fmt.Println("")
	return comando
}

func startMonitoring() {
	sites := readFile()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Iniciando monitoramento...")
		for _, site := range sites {
			testsSites(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testsSites(site string) {
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
func readFile() []string {
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

func writeFile(site string, status bool) {
	// define o arquivo para abrir, define que vai:
	// os.O_RDWR -> abre o arquivo leitura/gravação.
	// os.O_APPEND -> acrescenta dados ao arquivo ao gravar.
	// os.O_CREATE -> cria um novo arquivo se nenhum existir. 
	// 0666 -> codigo de permissão do arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(
		time.Now().Format("02/01/2006 15:04:05") + 
		" - " + site + "- online: " + 
		strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func showLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)
}
