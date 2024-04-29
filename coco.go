package main

import "fmt"

// func tiket(hari string, penumpang int)int{
// 	var hargaTiket int = 150000
// 	var diskon int
// 	if hari == "senin" && penumpang >= 3 {
// 			diskon = hargaTiket * 1/4
// 			hargaTiket = hargaTiket - diskon
// 	} else {
// 		return hargaTiket
// 	}
// 	return hargaTiket
// }

func properti(pavling int, totalBeli, totalJual *float64){
	var hargaBeli, hargaJual float64
	var luasTanah int
	for i := 1; i <= pavling; i++{
		fmt.Scan(&luasTanah, &hargaBeli, &hargaJual)
		hargaBeli = hargaBeli * float64(luasTanah)
		hargaJual = hargaJual * float64(luasTanah)
		*totalBeli += hargaBeli
		*totalJual += hargaJual
	}

}

func main(){
	var n int
	var nilaiBeli, nilaiJual float64
	fmt.Scan(&n)
	properti(n, &nilaiBeli, &nilaiJual)
	fmt.Println(nilaiBeli/1000000)
	fmt.Println(nilaiJual/1000000)
}