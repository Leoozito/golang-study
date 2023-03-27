// eu mesmo que fiz então está diferente -------------

package main

import (
	"fmt"
)

func main() {
	minhasgoroutines(10)
}

func minhasgoroutines(num int) {
	canal := make(chan int) 	// começando fazendo um canal aqui!
		go func() {
			for num := 1; num < 100; num++ {
				canal <- num
			}
			close(canal)
		}()

	for quant := 0; quant < 100; quant++ {
		fmt.Println(quant, "\t", <-canal)
	}
}
