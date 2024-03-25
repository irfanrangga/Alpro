package main

import "fmt"

func belajarMain(aktivitas string, belajar, main *int) {
	if aktivitas == "belajar" {
		*belajar += 1
	} else if aktivitas == "main" {
		*main += 1
	}
}

func main() {
	var (
		jumlahBelajar, jumlahMain int
		aktivitas                 string
	)

	for aktivitas != "tidur" {
		fmt.Scan(&aktivitas)
		belajarMain(aktivitas, &jumlahBelajar, &jumlahMain)
	}
	fmt.Println(jumlahBelajar, jumlahMain)
}
