// - Usando uma função anônima auto-executável
// - Usando buffer

package main 

import (
	"fmt"
)

func main() {
	// COM BUFFER
	c := make(chan int, 1)

		c <- 42

	fmt.Println(<-c)

	//COM FUNÇÃO ANONIMA 
	d := make(chan int, 1)  

	go func() {
		d <- 59
	}()

	fmt.Println(<-d)
}

