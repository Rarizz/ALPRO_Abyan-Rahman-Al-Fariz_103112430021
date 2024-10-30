package main

import "fmt"

func main () {

	var bilanganBulat int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 1 {
		fmt.Print("bukan positif")
	} else {
		fmt.Print("positif")
	}
}