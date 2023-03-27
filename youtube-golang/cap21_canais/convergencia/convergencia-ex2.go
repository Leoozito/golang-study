package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
// - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.

	canal :=  converge(trabalho("Primeiro"), trabalho("Segundo"))

	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}

}

func trabalho(s string) chan string {
	canal := make(chan string)
		go func(s string, c chan string) {
			for i := 0; ;i++ {
				c <- fmt.Sprintf("Função %v diz: %v",s , i)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
			}
		}(s, canal)
	return canal
}

func converge (x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <- x
		}
	}()

	go func() {
		for {
			novo <- <- y
		}
	}()
	return novo
}