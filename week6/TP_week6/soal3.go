package main

import "fmt"

type pegawai struct {
	nama                 string
	masaKerja, gajiPokok int
	besarBonus           float64
}

func main() {
	var x pegawai

	fmt.Scan(&x.nama, &x.gajiPokok, &x.masaKerja)
	hitungBonus(&x)
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.0f\n", x.nama, x.besarBonus)
}

func hitungBonus(p *pegawai) {
	if p.masaKerja >= 10 {
		p.besarBonus += float64(p.gajiPokok) * 1.5
	} else if p.masaKerja < 10 && p.masaKerja >= 5 {
		p.besarBonus += float64(p.gajiPokok) * 0.75
	} else if p.masaKerja < 5 {
		p.besarBonus += float64(p.gajiPokok) * 0.5
	}
}
