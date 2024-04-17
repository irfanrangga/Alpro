package main

import "fmt"

const Max int = 10
type tabInt [Max] int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), average(data, nData))
}

func baca(A *tabInt, n *int) {
		for i := 0; i < Max; i++ {
			var x int
			fmt.Scan(&x)
			if x  == 0 {
				return
			}
			A[*n] = x
			*n++
		}
}

func cetak(A tabInt, n int) {
	for j := 0; j < n; j++ {
		if A[j] > 0 {
			fmt.Print(A[j], " ")
		}
	}
	fmt.Println("")
}

func jumlah(A tabInt, n int)int {
	var hasil int
	for k := 0; k < n; k++ {
		hasil += A[k]
	}
	return hasil
}

func average(A tabInt, n int)float64 {
	return float64(jumlah(A, n))/float64(n)
}