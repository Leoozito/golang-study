package main

import (
	"fmt"
)

// - Chans par, ímpar, quit
// - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// - Func receive é um select entre os três canais, encerra no quit

func main() {

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)
	
	go mandaNúmeros(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNúmeros (par, impar chan int, quit chan bool) {
	total := 50
	for i := 0; i <  total; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true 
}

func receive (par, impar chan int, quit chan bool) {
	for {
		select {
			case v := <-par:
				fmt.Println("O número", v, "é par")
			case v := <-impar:
				fmt.Println("O número", v, "é impar")
			case v, ok := <-quit:
				if !ok {
					fmt.Println("Deu ruim!")
				} else {
					fmt.Println("Encerrando. Recebemos", v)
				}
				return
		}
	}
}