package main

import "fmt"

func jumlahGenapKelipatan(bilangan int){
	var hasil int
	for bilangan >= 0 {
		fmt.Scan(&bilangan)
		if bilangan%4 == 0{
			hasil += bilangan
		}
	}
	fmt.Println(hasil)
}

func main(){
	var n int
	jumlahGenapKelipatan(n)

}
	