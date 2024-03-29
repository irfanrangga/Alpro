package main

import "fmt"
func main(){
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gcd(A,B))
}

func gcd(a, b int)int{
	if b == 0 {
		return a 
	} else { // c = 0
		return gcd(b, a%b)
	}
	
}