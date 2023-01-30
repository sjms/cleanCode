package main

import (
	"testing"
)

func TestValidCPF(t *testing.T) {
	cpf := "111.444.777-35"
	isCPF := IsCPF(cpf)
	if true != isCPF {
		t.Errorf("expected a true returned %t", isCPF)
	}
}

func TestInvalidCPF(t *testing.T) {
	cpf := "123.456.789-10"
	isCPF := IsCPF(cpf)
	if false != isCPF {
		t.Errorf("expected a false returned %t", isCPF)
	}
}
