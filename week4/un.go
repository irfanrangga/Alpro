package main

import (
	"fmt"
)

// Fungsi untuk mencari nilai suku ke-n dari deret Aritmatika
func UnAritmatika(a, b, n int) int {
	if n == 1 {
		// Basis: jika n adalah 1, kembalikan suku pertama dari deret (a)
		return a
	}
	// Panggil rekursif dengan suku ke-n sebelumnya (n-1) dan tambahkan beda (b)
	return UnAritmatika(a, b, n-1) + b
}

func main() {
	// Masukkan nilai a (suku pertama), b (beda), dan n (suku ke-n) dari deret Aritmatika
	var a, b, n int
	fmt.Println("Masukkan nilai suku pertama (a):")
	fmt.Scanln(&a)
	fmt.Println("Masukkan nilai beda (b):")
	fmt.Scanln(&b)
	fmt.Println("Masukkan nilai suku ke-n (n):")
	fmt.Scanln(&n)

	// Panggil fungsi UnAritmatika untuk mencari nilai suku ke-n
	fmt.Printf("Nilai suku ke-%d dari deret aritmatika adalah %d\n", n, UnAritmatika(a, b, n))
}
