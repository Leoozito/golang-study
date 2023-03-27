package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go envia(canal)
	recebe(canal)
}

func envia(s chan<- int) {
	s <- 42 // colocando o valor 42 no canal
}

func recebe(r <-chan int) {				// com isso você está tirando do canal as informações
	fmt.Println("O valor recebido do canal:", <-r)
}