package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last: "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln("DEU RUIM")
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) { //precisa que retorne um []byte e um error
	bs, err := json.Marshal(a) // retornando o error e o slice of byte
	if err != nil { // se não der erro...
		return []byte{}, fmt.Errorf("é uma cilada bino!") // ...retorna o erro e o slice of byte
	}
	return bs, nil
}
