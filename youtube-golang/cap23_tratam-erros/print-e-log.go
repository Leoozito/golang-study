package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
// as funções com defer não rodam
func foo() {
	fmt.Println("When os.Exit() is called, deferred functins don't run")
}