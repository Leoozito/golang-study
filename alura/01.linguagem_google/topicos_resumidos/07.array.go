// Quando trabalhamos com GO, normalmente não usamos ARRAY's, e sim SLICES.
// pois o ARRAY ele tem que ter uma certa quantidade de armazenamento de dados, assim quando definimos

package main

import "fmt"

func main() {
	resp := lista()
	fmt.Println(resp)
}

func lista() [4]string {
	// criando a lista
	var lista [4]string

	// adicionando itens na lista
	lista[0] = "Leonardo"
	lista[1] = "Laís"
	lista[2] = "Vitoria"

	return lista
}
