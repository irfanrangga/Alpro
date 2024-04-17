package main

import "fmt"

func main() {
	var data tabInt
	fmt.Println("hasil", minimum(data, 5))
}
/*
func persamaan(x int, pers *int) {
	fmt.Scan(&x)
	*pers = 2 * x + 4
}

func cetakPointer(a int){
	var pers int
	persamaan(a, &pers)
	fmt.Println("Hasil: ", pers)
}
*/

/*
const NMax int = 1000

type tabT [NMax]int

func main() {
	var n, i int
	var DM, M, x, y, T tabT

	// input n
	fmt.Scan(&n)

	// input DM, M, x, y
	for i = 0; i < n; i++ {
		fmt.Scan(&DM[i], &M[i], &x[i], &y[i])
	}

	// Menampilkan Informasi banyaknya ramuan yang dapat di buat
	for i = 0; i < n; i++ {
		for DM[i] >= x[i] && M[i] >= y[i] {
			T[i]++
			DM[i] -= x[i]
			M[i] -= y[i]
		}
		fmt.Printf("total ramuan yang berhasil dibuat dari lemari persediaan %d sebanyak : %d\n", i+1, T[i])
	}
}
*/

//======================================================================================
/*
type lagu struct {
	judul  string
	artis  string
	durasi int
}

type playlist struct {
	laguList    [20]lagu
	totalDurasi int
	jumlahLagu  int
}

func insertLagu(pl *playlist, n int) {
	// I.S terdefinisi playlist kosong dan n lagu yang ingin dimasukkan.
	// F.S playlist terisi dengan n lagu
	for i := 0; i < n; i++ {
		fmt.Scanln(&pl.laguList[i].judul, &pl.laguList[i].artis, &pl.laguList[i].durasi)
		pl.totalDurasi += pl.laguList[i].durasi
	}
	pl.jumlahLagu = n
}

func isMember(pl playlist, la lagu) bool {
	// I.S Terdefinisi playlist dan sebuah lagu
	// F.S mengembalikan true jika lagu sudah termasuk dalam playlist
	for i := 0; i < pl.jumlahLagu; i++ {
		if pl.laguList[i].judul == la.judul && pl.laguList[i].artis == la.artis && pl.laguList[i].durasi == la.durasi {
			return true
		}
	}
	return false
}

func gabungPlaylist(pl1 playlist, pl2 playlist, pl3 *playlist) {
	// I.S Terdifinisi playlist zaki dan playlist mifta yang ingin digabung
	// F.S Playlist zaki dan playlist mifta tergabung pada playlist 3, Note : Playlist Zaki di bagian atas array.
	for i := 0; i < pl1.jumlahLagu; i++ {
		pl3.laguList[i] = pl1.laguList[i]
	}
	pl3.jumlahLagu = pl1.jumlahLagu

	for j := 0; j < pl2.jumlahLagu; j++ {
		if !isMember(*pl3, pl2.laguList[j]) {
			pl3.laguList[pl3.jumlahLagu] = pl2.laguList[j]
			pl3.jumlahLagu++
		}
	}
}

func tampilPlaylist(pl playlist) {
	// I.S  Terdefinisi playlist
	// F.S  judul dan artis semua lagu ditampilkan dan total durasi ditampilkan pada akhiran
	for i := 0; i < pl.jumlahLagu; i++ {
		fmt.Printf("%s %s\n", pl.laguList[i].judul, pl.laguList[i].artis)
	}
	fmt.Printf("Durasi Playlist: %d\n", pl.totalDurasi)
}

// Fungsi utama
func main() {
	var pl1, pl2, pl3 playlist
	var nZaki, nMifta int

	fmt.Scanln(&nZaki)
	insertLagu(&pl1, nZaki)
	fmt.Scanln(&nMifta)
	insertLagu(&pl2, nMifta)
	gabungPlaylist(pl1, pl2, &pl3)
	tampilPlaylist(pl3)
}
*/

/*
func isiArray(data ...int) [100]int {
	// mengembalikan array yang berisi bilangan bulat
	var arr [100]int
	for i := 0; i < len(data) && i < 100; i++ {
		arr[i] = data[i]
	}
	return arr
}

func countPositiveSumNegative(arr [100]int) (int, int) {
	// mengembalikan jumlah bilangan positif dan negatif dari array 
	var positive, negative int
	for i := 0; i < len(arr) && arr[i] != 0; i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative += arr[i]
		}
	}
	return positive, negative
}

func main() {
	// Input data
	var n int
	fmt.Scanln(&n)

	var data []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		data = append(data, num)
	}

	// Mengisi array dengan data input
	arr := isiArray(data...)

	// Menghitung jumlah bilangan positif dan negatif
	positive, negative := countPositiveSumNegative(arr)

	// Menampilkan hasil
	fmt.Printf("%d %d\n", positive, negative)
	
}
*/

const Nmax int = 5
type tabInt [Nmax] int

func minimum(T tabInt, n int)int {
	var min int = T[0]
	var k int

	for k = 1; k < Nmax; k++ {
		fmt.Scan(&T[k])
	}
 
	if k < n {
		if min > T[k] {
			min = T[k]
		}
		k++
	}
	return min
}