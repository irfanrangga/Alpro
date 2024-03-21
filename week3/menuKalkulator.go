package main

import "fmt"

func hitungJumlah(){
	var a, b, hasil int
	fmt.Print("Masukan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	hasil = a + b
	fmt.Println("Hasil penjumlahan: ", hasil)
}
func hitungPerkalian(){
	var a, b, hasil int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	hasil = a * b
	fmt.Println("Hasil perkalian: ", hasil)
}
func hitungPembagian(){
	var a, b, hasil int
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	hasil = a / b
	if b == 0 {
		fmt.Println("error: pembagi tidak boleh 0")
		fmt.Scan(&a, &b)
	} else {
		fmt.Println("Hasil Pembagian: ", hasil)
	}
}
func menu(){
	fmt.Println("--------------------")
	fmt.Printf("%13s\n", "M E N U")
	fmt.Println("--------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("--------------------")
}
func main(){
	var pilih int
	for {
		menu()
		fmt.Println("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch  {
		case pilih == 1:
			hitungJumlah()
		case pilih == 2:
			hitungPerkalian()
		case pilih == 3:
			hitungPembagian()
		}
		if pilih == 4{
			break
		}
	}
}