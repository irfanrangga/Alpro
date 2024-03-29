package main

import "fmt"

func main() {
	var a, b, c, hasil int
	var cek bool
	fmt.Println("Masukan bahan 1:")
	fmt.Scanln(&a)
	fmt.Println("Masukan bahan 2:")
	fmt.Scanln(&b)
	fmt.Println("Masukan bahan 3:")
	fmt.Scanln(&c)
	fmt.Println("Masukan total bahan:")
	fmt.Scan(&hasil)
	cek = cekHitung(a, b, c, hasil)
	if cek {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}
}

func cekHitung(b1, b2, b3 int, jumlah int) bool{
	if b1 + b2 + b3 == jumlah {
		return true
	} else {
		return false
	}
}