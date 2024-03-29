package main

import "fmt"

func main() {
	var a, b, total int
	fmt.Scan(&a, &b)
	hitungJumlah(a, b, &total)
	fmt.Println(total)
}

func hitungJumlah(b1, b2 int, h *int) {
	*h = b1 + b2
}
