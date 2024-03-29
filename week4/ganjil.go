package main

import "fmt"

func ganjil(n int)int{
	if n == 1 {
		return 1
	}else if n%2!=0{
		fmt.Print(n)
		return ganjil(n-2)	
	}
	return n
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(ganjil(N))
}