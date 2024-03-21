package main

func dinas(kelompok string, lama int) int {
	var A, B, C, total int
	A = 500000
	B = 400000
	C = 350000
	if lama < 2 {
		return 0
	}
	if kelompok == "A" || kelompok == "a" {
		total = A * (lama - 1)
	} else if kelompok == "B" || kelompok == "b" {
		total = B * (lama - 1)
	} else if kelompok == "C" || kelompok == "c" {
		total = C * (lama - 1)
	}
	return total
}
