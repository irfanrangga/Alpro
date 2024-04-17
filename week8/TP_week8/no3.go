package main

import "fmt"

const Nmax int = 20
type tabChar[Nmax] byte

func main() {
	var kata tabChar
	var nChar int
	baca(&kata, &nChar)
	cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
	fmt.Scanln(n)
	if *n > Nmax {
		*n = Nmax
	}

	for i := 0; i < *n; i++ {
		fmt.Scanf("%c", &k[i])
	}
}

func cetak(k tabChar, n int) {
	for j := n-1; j >= 0; j--{
		fmt.Printf("%c", k[j])
	}

}