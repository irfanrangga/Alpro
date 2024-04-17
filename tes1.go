package main

import "fmt"

// Tipe bentukan struct mahasiswa dengan field nama, NIM, dan nilai

type mahasiswa struct {
	nama string
	nim string
	nilai float64
}

func main() {
	var mhs1, mhs2, mhs3 mahasiswa
	// Panggil prosedur baca
	baca(&mhs1, &mhs2, &mhs3)
	// Panggil prosedur cetak
	cetak(mhs1, mhs2, mhs3)
	// Panggil fungsi rata_rata_nilai
	rata_rata_nilai(mhs1, mhs2, mhs3)
	// Panggil prosedur mhs_max_nilai
	mhs_max_nilai(mhs1, mhs2, mhs3)
}

func baca(m1, m2, m3 *mahasiswa) {
	/* IS: m1, m2, m3 terdefinisi sembarang
	   Proses: membaca dari piranti masukan data nama, NIM, dan nilai dari m1, m2, dan m3
	   FS: m1, m2, m3 berisi nilai
	*/
	fmt.Scan(&m1.nama, &m1.nim, &m1.nilai)
	fmt.Scan(&m2.nama, &m2.nim, &m2.nilai)
	fmt.Scan(&m3.nama, &m3.nim, &m3.nilai)
	
}

func cetak(m1, m2, m3 mahasiswa) {
	/* IS: m1, m2, m3 terdefinisi
	   FS: Semua field dari m1, m2, m3 tercetak di layar
	*/
	fmt.Println("")
	fmt.Println(m1.nama, m1.nim, m1.nilai)
	fmt.Println(m2.nama, m2.nim, m2.nilai)
	fmt.Println(m3.nama, m3.nim, m3.nilai)
}

func rata_rata_nilai(m1, m2, m3 mahasiswa) float64 {
    /* Mengembalikan rata-rata nilai dari 3 mahasiswa m1, m2, dan m3 */
	var average float64
	average = (m1.nilai + m2.nilai + m3.nilai) / 3
	fmt.Printf("Rata-rata nilai:%.2f\n", average)
	return average
}

func mhs_max_nilai(m1, m2, m3 mahasiswa) {
	/* IS: m1, m2, dan m3 terdefinisi
	   FS: Tercetak di layar nama mahasiswa dengan nilai tertinggi dengan format
	       "Mahasiswa dengan nilai tertinggi: <nama mahasiswa> <nilai mahasiswa>"
	       Asumsi nilai mahasiswa bersifat unik, sehingga hanya ada 1 nilai tertinggi
	*/
	if m1.nilai > m2.nilai && m1.nilai > m3.nilai {
		fmt.Printf("Mahasiswa dengan nilai tertinggi: %s %.2f", m1.nama, m1.nilai)
	} else if m2.nilai > m1.nilai && m2.nilai > m3.nilai {
		fmt.Printf("Mahasiswa dengan nilai tertinggi: %s %.2f", m2.nama, m2.nilai)
	} else if m3.nilai > m1.nilai && m3.nilai > m2.nilai {
		fmt.Printf("Mahasiswa dengan nilai tertinggi: %s %.2f", m3.nama, m3.nilai)
	}
	
}
