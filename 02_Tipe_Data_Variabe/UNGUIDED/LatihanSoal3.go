package main

import (
	"fmt"
)

func main() {

	var r float64

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	var opearasi float64 = 3.14 * r * r
	fmt.Println("Hasil dari luas lingkaran dengan jari-jari:", r, "adalah", opearasi)
}