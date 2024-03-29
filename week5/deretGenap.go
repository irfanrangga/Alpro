package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(jumlahDeretGenap(n))
}
func jumlahDeretGenap(x int) int {
	/*  fungsi mengembalikan hasil penjumlahan dari seluruh
	    bilangan genap positif yang kurang dari atau sama dengan n */
	if x == 0 {
		return 0
	} else if x%2 != 0 {
		return jumlahDeretGenap(x - 1)
	}
	return x + jumlahDeretGenap(x-1) // hint : gunakan hasil tambah dengan fungsi rekursif
}
