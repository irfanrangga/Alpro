package main

import "fmt"

type seriesReview struct {
	nama        string
	episode     int
	durasi      int //durasi per eps
	waktuTonton int
}

func main() {
	var series seriesReview
	var hasil, totalDurasi int
	fmt.Println("Masukkan judul series:")
	fmt.Scan(&series.nama)
	fmt.Println("Masukkan jumlah episode:")
	fmt.Scan(&series.episode)
	fmt.Println("Masukkan durasi per episode:")
	fmt.Scan(&series.durasi)
	fmt.Println("Masukkan berapa waktu tonton dalam sehari:")
	fmt.Scan(&series.waktuTonton)
	totalDurasi = series.episode * series.durasi
	hasil = (totalDurasi + (60 * series.waktuTonton) - 1) / (60 * series.waktuTonton)
	fmt.Print(hasil)
}
