package main

import "fmt"

const nMax int = 100

type belanja struct {
	nama  string
	harga int
}

type keranjang [nMax]belanja

func main() {
	var n int
	var arrKeranjang keranjang

	inputBelanja(&arrKeranjang, &n)
	tampilData(&arrKeranjang, &n)
}

func inputBelanja(arrKeranjang *keranjang, n *int) {
	fmt.Scanln(n)

	for i := 0; i < *n; i++ {
		fmt.Scanln(&arrKeranjang[i].nama, &arrKeranjang[i].harga)
	}
}

func tampilData(arrKeranjang *keranjang, n *int) {
	totalHarga := 0
	for i := 0; i < *n; i++ {
		fmt.Printf("%d. %s\n", i+1, arrKeranjang[i].nama)
		totalHarga += arrKeranjang[i].harga
	}
	fmt.Printf("Total belanja: %d\n", totalHarga)
}
