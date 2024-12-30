package main

import "fmt"

func main() {
	var x, jumlah int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&x)

	for x > 0 {
		digit := x%10
		fmt.Print(digit)
		jumlah += digit
		x /= 10
	}

	fmt.Println("")
	fmt.Println("Jumlah digit:", jumlah)
}
