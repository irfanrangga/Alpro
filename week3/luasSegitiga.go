package main

import "fmt"

func pecahUang(uang int, k10, k2, k1 *int) {
	*k10 = uang % 10
	*k2 = uang % 10 / 3000
	*k1 = uang % 1000
}

func main() {
	var uang, k10, k2, k1 int

	fmt.Scan(&uang)
	pecahUang(uang, &k10, &k2, &k1)
	fmt.Println(k10, k2, k1)
}
