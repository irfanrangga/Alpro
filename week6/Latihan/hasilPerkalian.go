package main

import "fmt"

type HasilPerkalian struct {
	totalProduk    float64
	jumlahBilangan int
}

func main() {
	hitungperkalian := perkalian()
	fmt.Println(hitungperkalian.totalProduk)
	fmt.Println(hitungperkalian.jumlahBilangan)
}

func perkalian() HasilPerkalian {
	hasilKali := HasilPerkalian{
		totalProduk:    1.0,
		jumlahBilangan: 0,
	}
	var bilangan int
	for {
		fmt.Scanln(&bilangan)
		if bilangan == 0 {
			break
		}

		if bilangan > 0 {
			hasilKali.totalProduk *= 1.0 / float64(bilangan)
			hasilKali.jumlahBilangan++
		}
	}
	return hasilKali
}
