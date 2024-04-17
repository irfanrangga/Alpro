package main

import "fmt"

const NMAX int = 5

type tabInt struct{
	info [NMAX] int
	n int
}

func main() {
	var nilaiAkhir tabInt
	inputNilai(&nilaiAkhir)
	inputNilai(&nilaiAkhir)
	inputNilai(&nilaiAkhir)
	inputNilai(&nilaiAkhir)
	inputNilai(&nilaiAkhir)
	inputNilai(&nilaiAkhir)
	outputNilai(nilaiAkhir)
}

func inputNilai(NA *tabInt) {
	/* NA.info[i] untuk menampung data nilai akhir, sedangkan
	NA.n untuk menampung banyaknya elemen data. Kedua field terdefinisi sembarang (bisa 0 bisa berisi) */
	if NA.n < NMAX {
		fmt.Scan(&NA.info[NA.n])
		NA.n++
	}
}

func outputNilai(NA tabInt) {
	//nilai akhir NA terdefinisi sembarang
	//F.S. seluruh elemn tercetak di layar
	for j:=0; j < NMAX; j++{
		fmt.Print(NA.info[j], " ")
	}
}