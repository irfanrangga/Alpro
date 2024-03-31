package main

import "fmt"

type titik struct {
	x, y  float64
	warna string
}

func main() {
	var x1, x2, y1, y2 float64
	var w1, w2 string
	var t1, t2 titik

	fmt.Scan(&x1, &y1, &w1)
	fmt.Scan(&x2, &y2, &w2)

	t1 = pembuatanTitikBaru(x1, y1, w1)
	t2 = pembuatanTitikBaru(x2, y2, w2)

	fmt.Printf("Data titik 1: koordinat (%v, %v), warna %s\n", t1.x, t1.y, t1.warna)
	fmt.Printf("Data titik 2: koordinat (%v, %v), warna %s\n", t2.x, t2.y, t2.warna)
}

func pembuatanTitikBaru(x, y float64, w string) titik {
	return titik{x, y, w}
}
