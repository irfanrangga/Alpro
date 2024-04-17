package main

import "fmt"

type tableName struct {
	firstName  string
	middleName string
	lastName   string
}

type tableAddress struct {
	street string
	city   string
	zip    string
}

type student struct {
	name    tableName
	address tableAddress
}

func main() {
	//var student1 student
	var nama student
	nama.name.firstName = "Irfan"
	nama.name.middleName = "Rangga"
	nama.name.lastName = "Miftahurrizqi"
	fmt.Println(nama.name.firstName)
	fmt.Println(nama.name.middleName)
	fmt.Println(nama.name.lastName)
	nama.address.city = "Cikarang"
	nama.address.street = "Bekasi"
}
