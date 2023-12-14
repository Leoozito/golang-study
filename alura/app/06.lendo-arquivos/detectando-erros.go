package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 2

func main () {
	exibeIntroducao()
	exibeMenu()
	//iniciarMonitoramento()
	// leSitesDoArquivo()
	// exibeNomes()

	// _, idade := devolveNomeeIdade()
	// fmt.Println("Tenho", idade, "anos")

	comando := leComando()
	for {
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
	sites := leSitesDoArquivo()
	
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

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		registraLog(site, true)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "esta com problemas. StatusCode:", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {
	var sites []string 

	arquivo, err := os.Open("sites.txt")

	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for i := 0; i < 6; i++ {
		
		linha, err := leitor.ReadString('\n') // tem q usar aspas simples para representar quebra de linha e não string
		linha = strings.TrimSpace(linha) 
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	// Para criar um arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //"0666" uma permissão padrão quando quiser usar essas funções do OS

	if err != nil {
		fmt.Println(err)
	}

	// que escreve pra gente no arquivo
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}