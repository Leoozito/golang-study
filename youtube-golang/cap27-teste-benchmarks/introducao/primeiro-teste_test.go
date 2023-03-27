package main

import (
	"testing"
)
// soma(i ...int) int

func TestSoma(t *testing.T) {   // caracteristica de um teste de uma função
	teste := soma(3,2,1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {   // caracteristica de um teste de uma função
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got:", teste)
	}
}