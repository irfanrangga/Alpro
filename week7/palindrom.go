package main

import (
	"fmt"
)
const maxSize = 256
type tabInt [maxSize] int

func isiArray(arr *tabInt, n int) {
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
	
}

func balikArray(arr *tabInt, n int) {
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
	
}

func palindrome(arr *tabInt, n int) bool {
    for i := 0; i < n/2; i++ {
        if arr[i] != arr[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var array tabInt
    var n int
    fmt.Print("Masukkan panjang array: ")
	fmt.Scan(&n)

	if n > maxSize {
        fmt.Println("melebihi kapasitas maksimal.")
        return
    }

    isiArray(&array, n)

    fmt.Println("sebelum dibalik:")
    fmt.Println(array[:n])

    balikArray(&array, n)

    fmt.Println("setelah dibalik:")
    fmt.Println(array[:n])


    if palindrome(&array, n) {
        fmt.Println("palindrom.")
    } else {
        fmt.Println("bukan palindrom.")
    }
}