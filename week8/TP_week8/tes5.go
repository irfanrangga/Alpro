package main

import (
	"fmt"
	"math"
)

const MAX int = 10

type tabInt [MAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(d *tabInt, n *int) {

	for i := 0; i < MAX; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			return
		}
		d[*n] = x
		*n++
	}
}

func cetak(d tabInt, n int) {
	for i := 0; i < n; i++ {
		if d[i] > 0 {
			fmt.Print(d[i], " ")
		}
	}
	fmt.Println()
}

func jumlah(d tabInt, n int) int {
	var jumlah int = 0
	for i := 0; i < n; i++ {
		jumlah += int(math.Abs(float64(d[i])))
	}
	return jumlah
}

func rata_rata(d tabInt, n int) float32 {
	return float32(jumlah(d, n)) / float32(n)
}