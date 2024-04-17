package main

import "fmt"

func main() {
	var bil, hasil int
	counter(bil, &hasil)
	fmt.Println("jumlah bilangan:", hasil)
}

func counter(n int, total *int){
	*total = 0
	for n = 1; n <= 100; n++ {
		switch  {
		case n%6 == 0:
			*total++
		case n%9 == 0:
			*total++
		case n%15 == 0:
			*total++
		case n%20 == 0:
			*total++
		}
	}
}