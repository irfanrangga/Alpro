package main

import "fmt"

type mobil struct {
	merek         string
	tahunProduksi int
	kecepatan     int
}

func main() {
	var m1, m2, m3 mobil
	var rataRata float64

	fmt.Scan(&m1.merek, &m1.tahunProduksi, &m1.kecepatan)
	fmt.Scan(&m2.merek, &m2.tahunProduksi, &m2.kecepatan)
	fmt.Scan(&m3.merek, &m3.tahunProduksi, &m3.kecepatan)

	rataRata = float64((m1.kecepatan + m2.kecepatan + m3.kecepatan) / 3)
	fmt.Printf("Rata-rata kecepatan mobil %s (%d), ", m1.merek, m1.tahunProduksi)
	fmt.Printf("Mobil %s (%d), dan mobil %s (%d): ", m2.merek, m2.tahunProduksi, m3.merek, m3.tahunProduksi)
	fmt.Printf("%.2f", rataRata)
}
