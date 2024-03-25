package main

import (
	"fmt"
	"math"
)

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = math.Pi * math.Pow(r, 2)
	*k = 2 * math.Pi * r
}
func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}
func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}
func main() {
	var r, s, luasL, luasP, kelL, kelP float64
	var totalLuas, totalKel float64

	fmt.Scan(&r, &s)
	for r != 0 && s != 0 {
		hitungLuasKelilingLingkaran(r, &luasL, &kelL)
		hitungLuasKelilingPersegi(s, &luasP, &kelP)
		hitungTotal(luasL, luasP, kelL, kelP, &totalLuas, &totalKel)
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, luasL, luasP, kelL, kelP, totalLuas, totalKel)
		fmt.Scan(&r, &s)

	}
}
