package main

import "fmt"

func lowerToUpper(kar rune) rune {
	if kar >= 'a' && kar <= 'z' {
		return kar - ('a' - 'A')
	}
	return kar
}
func main() {
	var kar rune
	fmt.Scanf("%c", &kar)
	fmt.Printf("%c", lowerToUpper(kar))
}

// z = 122
// 122  - ( 97 - 65 ) = 66
