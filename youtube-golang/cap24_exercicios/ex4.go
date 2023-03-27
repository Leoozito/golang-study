// use o struct sqrt.Error como valor do tipo erro

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erroNovo := fmt.Errorf("é uma cilada que ocorreu! %v", f)
		return 0, sqrtError{"123","321", erroNovo}
	}
	return 42, nil
}