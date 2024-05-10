package main

import "fmt"
type dataDesa struct{

}

type dataPenduduk struct{
	
}

func main() {
	menu()
}

func menu() {
	fmt.Println("================================")
	fmt.Printf("%18s\n","SI DESA")
	fmt.Println(" Aplikasi Sistem Informasi Desa")
	fmt.Println("================================")
	fmt.Println("Pilih menu:")
	fmt.Println("1. Input Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Ubah Data")
	fmt.Println("4. Exit")
	fmt.Println("================================")
	fmt.Println()
}

func inputData()