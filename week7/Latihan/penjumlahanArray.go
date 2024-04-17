package main

import "fmt"

const NMax int = 1000
type tabG [NMax]int

func main() {
	var N, i, hasil int
	var K string
	var KA, KB tabG

    //input banyaknya siswa
	fmt.Scan(&N)

    //input nilai kelompok A sebanyak N
	for i = 0; i < N; i++ {
	    fmt.Scan(&KA[i])
	}

    //input nilai kelompok B sebanyak N
	for i = 0; i < N; i++ {
		fmt.Scan(&KB[i])
	}

	//input kriteria yang ingin digunakan (ganjil/genap)
	fmt.Scan(&K)
	
	//Menjumlahkan elemen dari 2 array dengan kriteria masukan
	if K == "genap" {
		hasil = 0
		//menjumlahkan 2 array sesuai kriteria
		for i = 0; i < N; i++ {
			if KA[i] %2 == 0 {
				hasil = hasil + KA[i]
			}
			if KB[i]%2  == 0 {
				hasil = hasil + KB[i]
			}
		}
	} else if K == "ganjil" {
		hasil = 0
		//menjumlahkan 2 array sesuai kriteria
		for i = 0; i < N; i++ {
			if KA[i] %2 != 0 {
				hasil = hasil + KA[i]
			}
			if KB[i]%2 != 0 {
				hasil = hasil + KB[i]
			}
		}
	}
	fmt.Printf("Hasil dari penjumlahan elemen array yang bernilai %s adalah %d", K, hasil)
}