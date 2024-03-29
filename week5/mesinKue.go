package main

import "fmt"

func main() {
	var wheat, sugar, butter, egg, total int
	fmt.Scan(&wheat, &sugar, &butter, &egg)
	total = mesinKue(wheat, sugar, butter, egg)
	fmt.Println(total)
}

func mesinKue(terigu, gula, mentega, telur int) int {
	var jumKue int
	jumKue = terigu / 20
	if gula/5 < jumKue {
		jumKue = gula / 5
	} else if mentega/6 < jumKue {
		jumKue = mentega / 6
	} else if telur < jumKue {
		jumKue = telur
	}
	return jumKue
}
