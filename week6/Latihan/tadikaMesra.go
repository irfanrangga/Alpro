package main

import (
	"fmt"
)

// Definisikan tipe bentukan untuk data siswa
type Siswa struct {
	Nama        string
	JumlahSuara int
}

func main() {
	var calonKetua [3]Siswa
	for i := 0; i < 3; i++ {
		fmt.Scanln(&calonKetua[i].Nama)
	}

	// Input data voting kelas
	var totalSuara int
	for i := 0; i < 8; i++ {
		var pilihan int
		fmt.Scanln(&pilihan)
		calonKetua[pilihan-1].JumlahSuara++
		totalSuara++
	}

	// Menentukan pemenang
	var pemenang Siswa
	for _, calon := range calonKetua {
		if calon.JumlahSuara > pemenang.JumlahSuara {
			pemenang = calon
		}
	}

	// Menampilkan pemenang dan jumlah suara yang diperoleh
	fmt.Printf("%s\n", pemenang.Nama)
	fmt.Printf("%d", pemenang.JumlahSuara)
}
