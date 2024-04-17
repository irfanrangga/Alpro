package main

import "fmt"

const Nmax int = 1000
type tabInt [Nmax] int
type dataMahasiswa struct{
	Nama string
	Nim string
}

func main() {
	var (
		eprt tabInt
		ipk tabInt
		semester tabInt
		x int
		name, nim dataMahasiswa
	)

	fmt.Scan(&x)
	for i := 0; i < x ; i++{
			fmt.Print("Masukan Nama: ")
			fmt.Scan(&name.Nama)
			fmt.Print("Masukan NIM Mahasiswa: ")
			fmt.Scan(&nim.Nim)
			fmt.Print("Masukan nilai EPRT: ")
			fmt.Scan(&eprt[i])
			fmt.Print("Masukan IPK: ")
			fmt.Scan(&ipk[i])
			fmt.Print("Masukan semester: ")
			fmt.Scan(&semester[i])
			fmt.Println()
	}
	fmt.Printf("EPRT tertinggi: %d  IPK terendah: %d rata-rata semester lulusan: %.2f", highestEprt(&eprt, &x), lowestGpa(&ipk, &x), averageGrade(&semester, &x))
}

func highestEprt(a *tabInt, n *int)int {
	var k int
	var high int = a[0]

	if *n > Nmax {
		*n = Nmax
	}

	if k < *n {
		if high < a[k]{
			high = a[k]
		}
	}

	return high
}

func lowestGpa(b *tabInt, n *int)int {
	var k int
	var min int = b[0]

	if *n > Nmax {
		*n = Nmax
	}
	
	if k < *n {
		if min > b[k]{
			min = b[k]
		}
		k++
	}
	return min
}

func averageGrade(c *tabInt, n *int)float64 {
	var hasil int
	for i:=0; i < *n; i++{
		hasil += c[i]
	}
	rataRata := float64(hasil) / float64(*n)
	return rataRata
}