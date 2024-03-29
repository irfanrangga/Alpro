package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(totalTabungan(x, y, z))
}

func totalTabungan(x, y, z int) int {
	var hari, jumlahTabungan, tabunganAwal int
	for hari = 1; hari <= z; hari++ {
		if hari%2 == 0 || hari%3 == 0 {
			tabunganAwal = x
			jumlahTabungan = tabunganAwal + y
		}
	}
	return jumlahTabungan
}
