package main

import "fmt"

func main() {

	var digit int

	fmt.Print("Masukkan digit: ")
	fmt.Scan(&digit)

	for digit > 0 {
		angka := digit % 10
		digit /= 10
		fmt.Println(angka)
	}
}