package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata(N)
}

func cetak_rerata(n int) {
	hitungRerata(n, 0, 1)
}

func hitungRerata(n, sum, i int ) {
	var rerata float64
	if i > n {
		return
	}

	sum += i
	rerata = float64(sum) / float64(i)
	fmt.Println(sum, rerata)
	hitungRerata(n, sum, i+1)
}
