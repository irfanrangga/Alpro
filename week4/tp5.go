package main

import "fmt"

func sukuKeN(n int)int{
	if n <= 3 {
		return 1
	} else{
		return sukuKeN(n-1) + sukuKeN(n-2) + sukuKeN(n-3)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(sukuKeN(N))
}