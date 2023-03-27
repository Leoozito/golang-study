// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
package main

import "fmt"

// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("Deu zica! e olha isso: %v", e.qualquercoisa)
}

// - Crie uma função que tenha um valor do tipo error como parâmetro. 
func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).qualquercoisa, "e no metodo Error eu tenhp:", e)
}

func main() {
	minhaVariavel := erroEspecial{"Uma frase para a string qualquer coisa"}
	erroComoParametro(minhaVariavel)
}

