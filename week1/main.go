package main

import "fmt"

func main() {
	// --( NO.1 Soal Tabungan Kakak )--
	var tabAwal, lama int
	fmt.Print("Masukkan tabungan awal: ")
	fmt.Scan(&tabAwal)
	fmt.Print("Masukkan berapa lama hari: ")
	fmt.Scan(&lama)
	fmt.Println(uangKakak(tabAwal, lama))

	// --( NO.2 Soal Keliling Lingkaran )--
	//var jariJari float64
	//fmt.Print("Masukan jari-jari: ")
	//fmt.Scan(&jariJari)
	//fmt.Println(kelilingLingkaran(jariJari))

	//--( NO.3 Soal Persamaan )--
	//var x float64
	//fmt.Print("Masukan nilai x:")
	//fmt.Scan(&x)
	//fmt.Println(persamaan(x))

	// --( NO.4 Soal Dinas )--
	//var (
	//	kelompok string
	//	waktu    int
	//)
	//fmt.Print("Masukan kelompok: ")
	//fmt.Scan(&kelompok)
	//fmt.Print("Masukan berapa hari: ")
	//fmt.Scan(&waktu)
	//fmt.Println(dinas(kelompok, waktu))
}
