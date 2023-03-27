/* - Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race */

/* 
package main

import (
	"fmt"
	"runtime"
	"sync"
)
var cont = 1
var wg sync.WaitGroup
var mu sync.Mutex
// const quantidadeDeGourotines = 5

func main () {
	
	partegoroutine(1) // ou partegoiroutines(quantidadeDeGoroutines)
	wg.Wait()
	
}

func partegoroutine (s int) {
	wg.Add(s)
	for r := 1; r < s; r++ {
	go func() {
		mu.Lock() // padrão do MUTEX ser colocado no inicio da função
		minhaVariavel := cont
		runtime.Gosched()
		minhaVariavel++
		cont = minhaVariavel
		fmt.Println("Total de goroutines:\t", s, "\nTotal do contador:\t", cont)
		mu.Unlock() // e no final -> padrão MUTEX
		wg.Done() 
	}()
	}
} */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutines = 50000

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}