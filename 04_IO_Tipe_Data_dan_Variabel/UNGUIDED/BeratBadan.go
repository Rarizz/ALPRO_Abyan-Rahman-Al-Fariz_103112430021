package main

import "fmt"

func main() {
	var bmi float64
	var tinggiBadan float64
	var hasil float64

	fmt.Print("Masukkan nilai BMI anda: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi badan anda (m): ")
	fmt.Scan(&tinggiBadan)

	hasil= bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("Hasil berat badan anda adalah: %.0f", hasil)
}