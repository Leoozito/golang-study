package main

import (
	"fmt"
	"os"
)

func main() {

	// no GO não tem "WHILE", então utilizamos o FOR
	// e quando passamos o FOR sem nada, ele vira um loop infinito 
	for {
		exibeMenu()

		comando := retornaComando()

		switch comando {
		case 1:
			FOR_with_RANGE()
		case 2:
			fmt.Println("Monitorando...")
			FOR_default()
		case 3:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa !")
			os.Exit(0)
		default:
			fmt.Println("Comando invalido")
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Ver funcionamento do FOR com RANGE [adicional importante]")
	fmt.Println("2- Iniciar Monitoramento")
	fmt.Println("3- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func retornaComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

// FOR COM "RANGE"
func FOR_with_RANGE() {
	nomes := []string{"Leonardo", "Lais", "Vitoria"}
	fmt.Println("--------------------------------------------------")
	fmt.Println(" LOOP FOR COM RANGE ")
	fmt.Println(" ")
	fmt.Println("Essa é sua lista: ", nomes)

	for i, nome := range nomes { // o range do for é IMPORTANTE:
		// com o range, é possivel acessar as posições da lista ("i") e o valor das posições ("nome")
		fmt.Println("O nome: ", nome, "está na posição: ", i)
	}

	fmt.Println(" ")
	fmt.Println("--------------------------------------------------")
}

// FOR de lógica PADRÃO
func FOR_default() {
	nomes := []string{"Leonardo", "Lais", "Vitoria"}
	fmt.Println("--------------------------------------------------")
	fmt.Println(" LOOP FOR COM RANGE ")
	fmt.Println(" ")
	fmt.Println("Essa é sua lista: ", nomes)

    for i := 0; i < len(nomes); i++ {
        fmt.Println(nomes[i])
    }

	fmt.Println(" ")
	fmt.Println("--------------------------------------------------")
}