package main

import (
	"strconv"
	"strings"
)

const DV1 int = 10
const DV2 int = 11

func IsCPF(cpf string) bool {
	RemoveFormat(&cpf)
	isValid := IsValidLength(cpf)
	if !isValid {
		return false
	}
	doc := RemoveCheckDigits(cpf)
	doc += CalculateCheckDigit(doc, DV1)
	doc += CalculateCheckDigit(doc, DV2)
	return cpf == doc
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
