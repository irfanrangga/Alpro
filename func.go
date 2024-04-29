package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisGanjil(n)

}

func barisGanjil(n int){
	if n == 1 {
		fmt.Print(1)
	} else if n%2 == 1 {
		barisGanjil(n-1)
	}
	fmt.Print(n)
}