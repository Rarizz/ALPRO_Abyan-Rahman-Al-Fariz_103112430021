package main

import "fmt"

func main () {

	var bilanganBulat int
	var hasil bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 0 && bilanganBulat%2 ==0 {
		hasil = true
	}
	fmt.Print(hasil)
}