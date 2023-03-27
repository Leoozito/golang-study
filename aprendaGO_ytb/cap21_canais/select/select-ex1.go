package main

import (
	"fmt"
)

// - Duas go funcs enviando X/2 numeros cada uma pra um canal
// - For loop X valores, select case ←x

func main() {
	a := make(chan int)
	b := make(chan int)
	x:=50

	fmt.Println("")
	 
	go func(x int) {
		for y := 0 ; y < x; y++ {
			a <- y
		}
	}(x / 2)

	go func(x int) {
		for y := 0 ; y < x; y++ {
			b <- y
		}
	}(x / 2)

	for y := 0 ; y < x; y++ {
		select {
			case v := <- a:
				fmt.Println("Canal A:", v)
			case v := <- b:
				fmt.Println("Canal B:", v)
		}
	}
}