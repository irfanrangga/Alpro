package main

import "fmt"

func main() {
	var a, b, c int
	a, b, c = 0, 0, 0
	game(&a, &b, &c)
	fmt.Println(a, b, c)
}

func game(m1, m2, m3 *int) {
	var t1, t2, t3 int
	var p1, p2, p3 bool
	fmt.Scan(&t1, &t2, &t3)
	for t1 != 0 && t2 != 0 && t3 != 0 {
		if t1%2 == 0 {
			p1 = true
		} else if t1%2 != 0 {
			p1 = false
		}
		if t2%2 == 0 {
			p2 = true
		} else if t2%2 != 0 {
			p2 = false
		}
		if t3%2 == 0 {
			p3 = true
		} else if t3%2 != 0 {
			p3 = false
		}
		if p1 && p2 == false && p3 == false || p1 == false && p2 && p3 {
			*m1 += 1
		}
		if p1 == false && p2 && p3 == false || p1 && p2 == false && p3 {
			*m2 += 1
		}
		if p1 == false && p2 == false && p3 || p1 && p2 && p3 == false {
			*m3 += 1
		}
		fmt.Scan(&t1, &t2, &t3)
	}
}
