package main

import "fmt"

func main() {

	var totalBelanjaAwal int
	var diskon int

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanjaAwal)

	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&diskon)

	hasil := totalBelanjaAwal - (diskon* totalBelanjaAwal/100)

	fmt.Println("Total belanja setelah diskon: ", hasil)
}