package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é programa do exercicio de compilação cruzada. Foi compilado num linux/amd64, e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)
}