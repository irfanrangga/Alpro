package main

import "fmt"

type siswa struct {
	Nama        string
	JumlahSuara int
}

func main() {
	var pemenang siswa
	var calonKetua [3]siswa

	for i := 0; i < 3; i++ {
		fmt.Scanln(&calonKetua[i].Nama)
	}

	var totalSuara int
	for i := 0; i < 8; i++ {
		var pilihan int
		fmt.Scanln(&pilihan)
		calonKetua[pilihan-1].JumlahSuara++
		totalSuara++
	}

	for _, calon := range calonKetua {
		if calon.JumlahSuara > pemenang.JumlahSuara {
			pemenang = calon
		}
	}

	fmt.Printf("%s\n", pemenang.Nama)
	fmt.Printf("%d", pemenang.JumlahSuara)
}
