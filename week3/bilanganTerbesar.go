package main

import "fmt"

func cetakTerbesar(b1, b2, b3, b4 int){
	var bilanganTerbesar = b1
	if b2 > bilanganTerbesar {
		bilanganTerbesar = b2
	}
	if b3 > bilanganTerbesar {
		bilanganTerbesar = b3
	}
	if b4 > bilanganTerbesar {
		bilanganTerbesar = b4
	}
	fmt.Println(bilanganTerbesar)
}

func main() {
	var a,b,c,d int
	fmt.Scan(&a, &b, &c, &d)
	cetakTerbesar(a, b, c, d)
}