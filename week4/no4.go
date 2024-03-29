package main

import (
	"fmt"
)

func PenjumlahanDigit(N int) int {
	if N < 10 {
		return N
	}
	return N%10 + PenjumlahanDigit(N/10)
}

func main() {
	var N int
	fmt.Println("Masukkan bilangan bulat N:")
	fmt.Scanln(&N)
	result := PenjumlahanDigit(N)
	fmt.Println(result)
}
