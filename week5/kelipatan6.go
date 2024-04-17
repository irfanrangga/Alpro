package main

import "fmt"

func main() {
	kelipatan6(100)
	fmt.Println("")
	kelipatan9(100)
	fmt.Println("")
	kelipatan6dan9(100)
}

func kelipatan6(n int){
	for i := 6; i < n; i += 6 {
		fmt.Print(i, " ")
	}
}

func kelipatan9(n int) {
	for i := 9; i < n; i += 9 {
		fmt.Print(i, " ")
	}
}

func kelipatan6dan9(n int) {
	var i int
	/*

	for i := 6; i < n; i++ {
		if i%6 == 0 && i%9 == 0 {
			fmt.Print(i, " ")
		}
	}

	*/
	
	if i >= n {
		fmt.Print(".")
	} else {
		if i%6 == 0 && i%9 == 0 {
			fmt.Print(i, " ")
		}
	}
}