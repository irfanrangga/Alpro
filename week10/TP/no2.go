package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX] int

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(___, ___)
	// cetak data pemain
	cetakPemain(___, ___)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(___, ___)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(___, ___)
}

func bacaData(A *tabPemain, n *int) {
	var data pemain
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&)
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	__
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	___
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	___
}
