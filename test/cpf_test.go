package test

import (
	"cleanCode/service"
	"testing"
)

func TestValidCPF(t *testing.T) {
	cpf := "111.444.777-35"
	isCPF, err := service.IsCPF(cpf)
	if true != isCPF || err != nil {
		t.Errorf("Expected a valid CPF returned, %t, %s", isCPF, err)
	}
}

func TestInvalidCPF(t *testing.T) {
	cpf := "123.456.789-10"
	isCPF, err := service.IsCPF(cpf)
	if true == isCPF {
		t.Errorf("Expected a invalid CPF returned, %t, %s", isCPF, err)
	}
}

func TestInvalidCPFWithRepeatedCharacters(t *testing.T) {
	cpf := "11111111111"
	isCPF, err := service.IsCPF(cpf)
	if true == isCPF {
		t.Errorf("Expected a invalid CPF returned, %t, %s", isCPF, err)
	}
}

func TestCPFWithInvalidLength(t *testing.T) {
	cpf := "123123123123"
	isCPF, err := service.IsCPF(cpf)
	if true == isCPF {
		t.Errorf("Expected a invalid CPF returned, %t, %s", isCPF, err)
	}
}

func TestEmptyCPF(t *testing.T) {
	cpf := ""
	isCPF, err := service.IsCPF(cpf)
	if true == isCPF {
		t.Errorf("Expected a invalid CPF returned, %t, %s", isCPF, err)
	}
}
