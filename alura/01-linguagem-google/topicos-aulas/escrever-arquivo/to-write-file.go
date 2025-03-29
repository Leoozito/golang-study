package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	registraLog("https://github.com/", true)	
}

func registraLog(site string, status bool) {

	// a função do pacote os, "OpenFile" requer:
	// nome do arquivo
	// "flags" - o que a função poderá fazer
	// e a permissões do arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if (err != nil) {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}