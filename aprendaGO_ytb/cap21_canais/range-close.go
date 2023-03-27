package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go meuloop(10, c)
	prints(c)
}

func meuloop(t int, s chan<- int) {
	for d:=0; d < t; d++ {
		s <- d //enviando o resultado "d" para o canal
	}
	close(s) //close serve para terminar a função executavel
}

func prints(r <-chan int) {
	for v := range c { //range fica diferente para canais
		fmt.Println("Recebido do canal:", v)
	}
}