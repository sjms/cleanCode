package service

import (
	"errors"
	"strconv"
	"strings"
)

const (
	DV1 int = 10
	DV2 int = 11
)

func IsCPF(cpf string) (bool, error) {
	RemoveFormat(&cpf)
	err := checkCPFIntegrity(cpf)
	if err != nil {
		return false, err
	}
	doc := RemoveCheckDigits(cpf)
	doc += CalculateCheckDigit(doc, DV1)
	doc += CalculateCheckDigit(doc, DV2)
	if cpf == doc {
		return true, nil
	}
	return false, errors.New("CPF inválido")
}

func CalculateCheckDigit(cpf string, index int) string {
	var sum int
	for _, r := range cpf {
		value := ToIntValue(r) * index
		sum += value
		index--
	}
	sum %= 11
	if sum < 2 {
		return "0"
	}
	return strconv.Itoa(11 - sum)
}

func IsValidLength(cpf string) bool {
	if len(cpf) == 11 {
		return true
	}
	return false
}

func HasRepeatedCharacters(cpf string) bool {
	repeated := strings.Repeat(cpf[:1], 11)
	return cpf == repeated
}

func checkCPFIntegrity(cpf string) error {
	if !IsValidLength(cpf) {
		return errors.New("o CPF possui tamanho inválido")
	}
	if HasRepeatedCharacters(cpf) {
		return errors.New("um CPF não pode possuir caracteres repetidos")
	}
	return nil
}

func RemoveFormat(cpf *string) {
	replacer := strings.NewReplacer(".", "", "-", "", " ", "")
	*cpf = replacer.Replace(*cpf)
}

func RemoveCheckDigits(cpf string) string {
	return cpf[:len(cpf)-2]
}

func ToIntValue(r rune) int {
	return int(r - '0')
}
