/* criar 3 goroutine (2 normal e 1 principal) 
   cada goroutine adicional deve fazer um print separado
   utiliza waitgroup para as pausas  */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasgoroutines(2)
	wg.Wait()      // como tem o "novasgoroutines" dentro de uma func, tem q dar o wg.wait para "esperar" para não acabar nessa função."
}

func novasgoroutines (i int) {
	wg.Add(i)
	for j := 0; j <  i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou goroutine número:", i+1)
			wg.Done() // PRONTO !
		} (x)
	}
}