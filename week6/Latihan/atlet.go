package main

import (
	"fmt"
)

type Atlet struct {
	AsalNegara string
	Waktu      float64
}

func main() {
	var atlet [3]Atlet
	for i := 0; i < 3; i++ {
		fmt.Scanln(&atlet[i].AsalNegara, &atlet[i].Waktu)
	}

	pemenang := cariPemenang(atlet)

	if pemenang != "" {
		fmt.Printf("Atlet asal %s menang", pemenang)
	} else {
		fmt.Println("Seri")
	}
}

func cariPemenang(atlet [3]Atlet) string {
	if atlet[0].Waktu < atlet[1].Waktu && atlet[0].Waktu < atlet[2].Waktu {
		return atlet[0].AsalNegara
	} else if atlet[1].Waktu < atlet[0].Waktu && atlet[1].Waktu < atlet[2].Waktu {
		return atlet[1].AsalNegara
	} else if atlet[2].Waktu < atlet[0].Waktu && atlet[2].Waktu < atlet[1].Waktu {
		return atlet[2].AsalNegara
	}
	return ""
}
