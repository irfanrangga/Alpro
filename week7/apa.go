package main

import "fmt"

const NMax int = 1000

type tabInt [NMax]int
type tabStr [NMax]string

func main() {
	var nama tabStr
	var jam, menit, detik, total tabInt
	var n, i, min, idx int

	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		fmt.Scanln(&nama[i], &jam[i], &menit[i], &detik[i])
	}

	for i = 0; i < n; i++ {
		total[i] = jam[i]*3600 + menit[i]*60 + detik[i]
	}

	min = total[0]
	idx = 0
	for i = 1; i < n; i++ {
		if total[i] < min {
			min = total[i]
			idx = i
		}
	}

	fmt.Printf("Peserta tercepat adalah %s dengan waktu %d jam %d menit %d detik \n", nama[idx], jam[idx], menit[idx], detik[idx])
}
