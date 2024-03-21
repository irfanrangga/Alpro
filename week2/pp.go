package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64 {
	var x float64
	x = T * math.Pi / 180
	return R * math.Cos(x)
}

func panjangY(R, T float64) float64 {
	var panjang float64
	var y float64
	y = T * math.Pi / 180
	panjang = R * math.Sin(y)
	return panjang
}

func main() {
	var R, T float64
	fmt.Scanf("%f %f", &R, &T)
	fmt.Printf("%.2f %.2f", panjangX(R, T), panjangY(R, T))
}
