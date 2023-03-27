// Usando select statement pra demonstrar os valores

package main 

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(q <-chan int, c chan int) {
	for {
		select {
			case v := <-c
				fmt.Println(v)
			case <- q:
				return
		}
	}
}