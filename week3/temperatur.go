package main

import "fmt"

func temperatur(celcius float32, reamur, fahrenheit, kelvin *float32) {
	*reamur = (4.0 / 5.0) * celcius
	*fahrenheit = celcius*9.0/5.0 + 32
	*kelvin = celcius + 273.15
}
func main() {
	var c, r, f, k float32
	fmt.Scan(&c)
	temperatur(c, &r, &f, &k)
	fmt.Printf("%.2fR %.2fF %.2fK", r, f, k)
}
