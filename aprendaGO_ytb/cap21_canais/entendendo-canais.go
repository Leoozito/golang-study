package main

import (
	"fmt"
)

func main() {
	// essa é a declaração para ter um canal
	meuCanal := make(chan int, 5)

	meuCanal <- 42 // essa goroutine está ** passando ** o valor 

	fmt.Println(<- meuCanal) // e essa goroutine (no caso a func main que também é uma goroutine) está ** pegando ** o valor
}