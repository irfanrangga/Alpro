package main

import "fmt"

func reamur(c float64) float64 {
	return c * (4.0 / 5.0)
}
func fahrenheit(c float64) float64 {
	return c*(9.0/5.0) + 32.0
}
func kelvin(c float64) float64 {
	return c + 273.0

}

func main() {
	var celciusAwal, n, step float64

	fmt.Print("Masukan suhu awal, suhu akhir, dan longkap:")
	fmt.Scan(&celciusAwal, &n, &step)
	fmt.Printf("%10.2s %10.2s %10.2s %10.2s\n", "C", "R", "F", "K")

	for i := celciusAwal; i <= n; i += step {
		R := reamur(i)
		F := fahrenheit(i)
		K := kelvin(i)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", i, R, F, K)
	}
}
