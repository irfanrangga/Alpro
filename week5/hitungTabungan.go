package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(totalTabungan(x, y, z))
}

func totalTabungan(x, y, z int) int {
	var hari, jumlahTabungan, tabunganAwal int
	tabunganAwal = x
	for hari = 1; hari <= z; hari++ {
		if hari%2 == 0 || hari%3 == 0 {
			jumlahTabungan += tabunganAwal
			tabunganAwal += y
		}
	}
	return jumlahTabungan
}
