package main

import "fmt"

func main() {
	var nilai int = 0
	for i := 1; i < 1000000000000; i++ {
		if i%6 == 0 || i%9 == 0 || i%15 == 0 || i%20 == 0 {
			nilai += 1
		}
	}
	fmt.Println(nilai)
}

