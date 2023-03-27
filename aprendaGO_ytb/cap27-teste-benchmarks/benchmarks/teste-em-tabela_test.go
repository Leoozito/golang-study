package main

import (
	"testing"
	"fmt"
)
// soma(i ...int) int
type test struct {
	data []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test {
		test{data: []int{1, 2, 3,}, answer: 6}, //answer seria a resposta
		test{[]int{10, 11, 12}, 33}, 
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := Soma(v.data ...)
		if z != v.answer {
			t.Error("Expected:", v.answer,"Got:", z)
		}
	}
}

func ExampleSoma () {
	fmt.Println(Soma(3, 2, 1))
}

func TestSoma(t *testing.T) {   // caracteristica de um teste de uma função
	teste := Soma(3,2,1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {   // caracteristica de um teste de uma função
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got:", teste)
	}
}

func BenchmarkSoma (b *testing.B) {
	for i := 0; i < b.N; i++ { 
		Soma(1, 1)
	}
}

func BenchmarkMultiplica (b *testing.B) {
	for i := 0; i < b.N; i++ { 
		Multiplica(2, 1)
	}
}