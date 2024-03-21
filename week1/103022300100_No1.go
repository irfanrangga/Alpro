package main

func uangKakak(tabPertama, lama int) int {
	var jumlah, i int
	jumlah = 0
	for i = 1; i <= lama; i++ {
		jumlah += tabPertama + (lama-i)*2500
	}
	return jumlah
}
