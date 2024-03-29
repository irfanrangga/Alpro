package main

import "fmt"

func cetakGoLang(i, n int) {
	/*  I.S bilangan bulat i dan n
		F.S menampilkan bilangan bulat i, string "Go", "Lang", atau "GoLang" */
		if i == n {
			if i%2 == 0 && i%3 != 0 {
				fmt.Println("Go")
			} else if i%3 == 0 && i%2 != 0 {
				fmt.Println("Lang")
			} else if i%6 == 0 {
				fmt.Println("GoLang")
			} else if i%2 != 0 && i%3 != 0 {
				fmt.Println(n)
			}
		} else if i%2 == 0 && i%3 != 0 {
			fmt.Println("Go")
			cetakGoLang(i+1, n)
		} else if i%3 == 0 && i%2 != 0 {
			fmt.Println("Lang")
			cetakGoLang(i+1, n)
		} else if i%6 == 0 {
			fmt.Println("GoLang")
			cetakGoLang(i+1, n)
		} else if i%2 != 0 && i%3 != 0 {
			fmt.Println(i)
			cetakGoLang(i+1, n) // gunakan fungsi rekursif di baris ini
		}
	}
func main() {
	var n int
	fmt.Scan(&n)
	cetakGoLang(1, n)
}