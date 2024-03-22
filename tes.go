package main

import (
	"fmt"
)

func kalkulator(operand1, operand2 float64, operator rune) float64 {
	var hasil float64
	switch {
	case operator == '+':
		hasil = operand1 + operand2
	case operator == '-':
		hasil = operand1 - operand2
	case operator == '*':
		hasil = operand1 * operand2
	case operator == '/':
		if operand2 != 0 {
			hasil = operand1 / operand2
		} else {
			return 0
		}
	}
	return hasil
}

func main() {
	var operand1, operand2 float64
	var operator rune

	fmt.Println("Masukkan 2 bilangan real dan operator (+, -, *, /): ")
	fmt.Scan(&operand1, &operand2, &operator)

	result := kalkulator(operand1, operand2, operator)
	fmt.Printf("Hasil: %T %T %T", operand1, operand2, result)
}
