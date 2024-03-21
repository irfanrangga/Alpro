package main

import "fmt"

func simpleCalc(b int, c *int) {
	var _ int
	_ = 10 + b - *c
	*c += 4
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	simpleCalc(b, &c)
	fmt.Println(a, b, c)
	fmt.Scan(&b)
	simpleCalc(c, &b)
	fmt.Println(a, b, c)
}
