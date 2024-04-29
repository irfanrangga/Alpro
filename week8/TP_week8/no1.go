package main

import "fmt"

const NMax int = 10
type tim struct{
	gol [NMAX] int
}

func main() {
	var nilaiAkhir tabInt
	var banyakNilai int
	bacaNilai(&nilaiAkhir,&banyakNilai)
	cetakNilai(nilaiAkhir, banyakNilai)
}

func bacaNilai(NA *tabInt, n *int, te *tim){
	fmt.Scan(n)
	if *n > NMax{
		*n = NMax
	}
	for i:= 0; i < *n; i++ {
		fmt.Scan(&NA[i])
		fmt.Scan(&te.gol[i])
	}
}

func cetakNilai(NA tabInt, n int){
	if n == 0{
		fmt.Println("Array kosong")
	} else {
		fmt.Print("Nilai yang terdapat pada array: ")
		for j := 0; j < n; j++ {
			fmt.Print(NA[j]," ")
		}
		fmt.Println("\n")
	}
}