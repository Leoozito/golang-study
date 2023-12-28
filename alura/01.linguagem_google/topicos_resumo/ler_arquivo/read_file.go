package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var resp int
	fmt.Println("Ver o funcionamento da biblioteca")
	fmt.Println("1. Ioutil")
	fmt.Println("2. Bufio")
	fmt.Scan(&resp)

	if resp == 1 {
		retorna_nomes_file()
	} else {
		retorna_alguns_nomes()
	}
}

func retorna_nomes_file() []string {
	var nomes []string
	arquivo, _ := ioutil.ReadFile("nomes.txt")

	fmt.Println(string(arquivo)) // o "ioutil" retorna bytes, e convertendo para string é possivel ver o que contém dentro do arquivo
	return nomes
}

func retorna_alguns_nomes() {
	arquivo, err := os.Open("nomes.txt") // abre o arquivo

	leitor := bufio.NewReader(arquivo) // Lê o arquivo completo em byte

	resp, err := leitor.ReadString('\n') // colocando esse formato, estou dizendo que é para o GO ler até uma determinada parte do arquivo, ou seja, ler até a "quebra de linha" (\n : significa quebra de linha)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(resp)
}
