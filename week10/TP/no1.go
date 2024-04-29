package main

import "fmt"

const Nmax int = 20
type tabInt [Nmax] int

func main() {
	var x tabInt
	var n int
	
	baca(&x, &n)
	cetakElemen(x, n)
	cetakInfo(x, n)
}

func baca(A *tabInt,n *int){
	for i := 0; i < Nmax; i++ {
		fmt.Scan(&A[i])
		if A[i] < 0 || A[i] == 0 {
			break
		}
		*n++
	}
}

func cetakElemen(A tabInt, n int){
	fmt.Print("Elemen array: ")
	for j := 0; j < n; j++{
		fmt.Print(A[j], " ")
	}
}

func maksimum(A tabInt, n int)int{
	var max int = A[0]

	for i := 0; i < n; i++{
		if max < A[i] {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int)int{
	var min int= A[0]

	for i := 0; i < n; i++ {
		if min > A[i] {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	fmt.Printf("\nNilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	fmt.Printf("Banyak elemen: %d\n", n)
}