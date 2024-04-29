package main

import "fmt"

const NMAX int = 5

type matrix [NMAX][NMAX]int

func main() {
	var m1, m2, m3 matrix
	var n int

	fmt.Scan(&n)                    // ukuran matriks

	fmt.Println("Matriks pertama")
	baca(&m1, &n)
	cetak(m1, n)

	fmt.Println("Matriks kedua")
	baca(&m2, &n)
	cetak(m2, n)

	fmt.Println("Matriks ketiga")
	bagi(m1,m2,&m3, n)
	cetak(m3, n)
}

func baca(m *matrix, n *int) {
	/*
		IS: Matriks m terdefinisi sembarang, n terdefinisi
		FS: Matriks m dengan baris n x kolom n berisi nilai. Jika n > NMAX
		    gunakan nilai n = NMAX
	 */
	if *n > NMAX {
		*n = NMAX
	} 
	for i := 0; i < *n; i++ {
		fmt.Scan(&m[i][i])
	}
}

func cetak(m matrix, n int) {
	/*
		IS: Matrik m terdefinisi, n terdefinisi
		FS: Tercetak matriks m di layar dengan format:
	
			x11 x12 ... x1n
			x21 x22 ... x2n
			... ... ... ...
			xn1 xn2 ...	xnn
	
			dengan n adalah ukuran matriks
	 */
	for j := 0; j < n; j++{
		fmt.Print(m[j][j])
	}
}

func bagi(A, B matrix, C *matrix, n int) {
	/*
		IS: Matriks A dan B terdefinisi
		FS: Matriks C berisi nilai. Elemen ke-i dan j matriks C merupakan 
		    pembagian bilangan bulat (div) elemen ke-i dan j dari matriks A oleh B.
			
			Contoh matriks dengan n = 2:
			matriks A div matriks B = matriks C
			4 5		 div   3 2	  =   1 2
			3 10		   4 3		  0 3
	 */
	for i := 0; i < n; i++{
		*C = A[i] / B[i]
	}
}
