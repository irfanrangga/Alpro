package main

import "fmt"
func main(){
	var N, bilangan int
	fmt.Scan(&N)
	cetakDeret(N, bilangan)
}

func cetakDeret(x, bil int){
	hitungFibo(x,0)
}

//  0 1 2 3 4 5 6
//  0 1 1 2 3 5 8
func hitungFibo(n,i int) int{
	fmt.Print(i)
	i += 1
	if n < i{
		return n
	}
	
	if n == 0{
		return 0
	} else if n <= 2 {
			return 1
	} else {
		//fmt.Print(hitungFibo(n, i))
		return hitungFibo(n-1, i) + hitungFibo(n-2, i)
	}
}
 func sukuKe(s,i int)int{
	if s < i {
		return 0
	}
	return s
 }