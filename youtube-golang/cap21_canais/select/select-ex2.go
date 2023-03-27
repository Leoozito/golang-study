package main

import (
	"fmt"
)

// - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
// - Func 2 for infinito, select: case envia pra canal, case recebe de quit

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0 // canal quit
}

func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
			case canal <- qualquercoisa: // select que envia coisa pro canal
				qualquercoisa++
			case <- quit: // e recebe coisa do canal
					return
		}
	}
}