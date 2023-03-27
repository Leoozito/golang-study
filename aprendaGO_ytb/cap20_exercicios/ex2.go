/* Esse exercicio vai eforçar seus conhecimentos sobre conjuntos de metodos
	Crie um tipo para um struct chamado "pessoa"[Crie um metodo "falar para este tipo que tenha um receiver ponteiro (*pessoa")
	Crie uma interface, "humanos", que seja implementada por tipos com o metodo "falar"
	 - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
 
package main

import (
	"fmt"
)

type pessoa struct {
	Nome string  
	Sobrenome string
	Idade int
}

type humanos interface {  // como é a interface
	falar()
}

func main () {
	p1 := pessoa{"Rhaissa", "Garcia", 18}

	p1.falar()

	dizerAlgumaCoisa(&p1)
}
// para criar um metodo necessita criar uma função
func (p *pessoa) falar() {
	fmt.Println(p.Nome, "diz oi!")
}
					// parametro
func dizerAlgumaCoisa (h humanos) {
	h.falar()
}