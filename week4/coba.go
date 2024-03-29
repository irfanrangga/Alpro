package main

import "fmt"

func faktorialLoop(n int, hasil *int) {

	var i int
	*hasil = 1
	for i = n; i >= 1; i-- {
		*hasil = *hasil * i
	}
}
func faktorialRecursive(n int) int {
	var hasil int
	if n == 1 {
		return 1
	} else {
		hasil = n * faktorialRecursive(n-1)
		return hasil
	}
}

func faktorialTailReccurcion(n, i int) int {
	if n <= 1 {
		return i
	}
	return faktorialTailReccurcion(n-1, n*i)
}

/*
faktorialTailReccurcion(5, 5*1)
faktorialTailReccurcion(4, 4*5)
faktorialTailReccurcion(3, 3*20)
faktorialTailReccurcion(2, 2*60)
faktorialTailReccurcion(1, 1*120)
n == 1
	return 120
*/

func main() {
	var faktorial int
	println("========= loop ========")
	faktorialLoop(5, &faktorial)
	fmt.Println(faktorial)
	println("========= Konventional ========")
	fmt.Println(5 * 4 * 3 * 2 * 1)
	println("========= rekursif ========")
	fmt.Println(faktorialRecursive(5))
	println("========= rekursif tail end recursion ========")
	fmt.Println((faktorialTailReccurcion(5, 1)))
}