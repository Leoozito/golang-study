package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	resp := menu()

	if resp == 1 {
		retorna_nomes_file()
	} else {
		retorna_alguns_nomes()
	}
}

func menu() int {
	var resp int
	fmt.Println("Ver o funcionamento da biblioteca")
	fmt.Println("1. Ioutil")
	fmt.Println("2. Bufio")
	fmt.Scan(&resp)
	return resp
}

// modo para o arquivo todo: usando a biblioteca "IOUTIL"
func retorna_nomes_file() []string {
	var nomes []string
	arquivo, _ := ioutil.ReadFile("nomes.txt") // usando a biblioteca IOUTIL

	fmt.Println(string(arquivo)) // o "ioutil" retorna bytes, e convertendo para string é possivel ver o que contém dentro do arquivo
	return nomes
}

// modo de acessar partes especificas do arquivo: com a biblioteca "BUFIO"
func retorna_alguns_nomes() {
	arquivo, err := os.Open("nomes.txt") // abre o arquivo

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo) // Lê o arquivo completo em byte

	for {
		resp, err := leitor.ReadString('\n') // colocando esse formato, estou dizendo que: é para o GO ler até uma determinada parte do arquivo, ou seja, ler até a "quebra de linha" ( \n : significa quebra de linha)
		
		resp = strings.TrimSpace(resp)
		
		if err == io.EOF { // EOF "End of File" - Componente da biblioteca bufio , que indica fim do arquivo
			break
		}
	}

	arquivo.Close()
}
