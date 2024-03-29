package main

import "fmt"

func kakiHewan(N int)int{
	var result int 
	if N == 1 {
		return 2
	}
	if N%2 == 0{
		result = kakiHewan(N-1) + 3
	} else {
		result = kakiHewan(N-1) + 2
	}
	return result
}
func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(kakiHewan(N))
}