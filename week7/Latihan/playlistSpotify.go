package main
import "fmt"

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
//  I.S terdefinisi playlist kosong dan n lagu yang ingin di masukkan.
//  F.S playlist terisi dengan n lagu
    ...
}

func isMember(pl playlist, la lagu) bool {
// I.S Terdefinisi playlist dan sebuah lagu
// F.S mengembalikan true jika lagu sudah termasuk dalam playlist
// note : Gunakan function ismember untuk mempermudah pergabungan
	...
}

func gabungPlaylist(pl1 playlist, pl2 playlist, pl3 *playlist) {
//  I.S Terdifinisi playlist zaki dan playlist mifta yang ingin digabung
//  F.S Playlist zaki dan playlist mifta tergabung pada playlist 3, Note : Playlist Zaki di bagian atas array.
    ...
}

func tampilPlaylist(pl playlist) {
//  I.S  Terdefinisi playlist
//  F.S  judul dan artis semua lagi ditampilkan dan total durasi ditampilkan pada akhiran
    ...
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
