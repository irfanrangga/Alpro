package main

import (
	"fmt"
	"math"
)

func hitungLuasKelilingLingkaran(r float64, l *float64, k *float64) {
	*l = math.Pi * r * r
	*k = 2 * math.Pi * r
}

func hitungLuasKelilingPersegi(s float64, l *float64, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(l1, l2, k1, k2 float64, totLuas *float64, totKel *float64) {
	*totLuas = l1 + l2
	*totKel = k1 + k2
}

func main() {
	var r, s, lLingkaran, lPersegi, kLingkaran, kPersegi, totLuas, totKel float64
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "L_U", "L_P", "K_L", "K_P", "T_L", "T_K")

	for {
		fmt.Scanf("%f %f", &r, &s)
		if r == 0 && s == 0 {
			break
		}
		hitungLuasKelilingLingkaran(r, &lLingkaran, &kLingkaran)
		hitungLuasKelilingPersegi(s, &lPersegi, &kPersegi)
		hitungTotal(lLingkaran, lPersegi, kLingkaran, kPersegi, &totLuas, &totKel)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, lLingkaran, lPersegi, kLingkaran, kPersegi, totLuas, totKel)
	}
}
