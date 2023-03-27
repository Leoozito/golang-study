package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go colocarlá(c)
	for v := range c {
		fmt.Println(v)
	}
}

func colocarlá(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}