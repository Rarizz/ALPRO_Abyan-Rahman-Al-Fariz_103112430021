package main

import "fmt"

func main(){
	var beratBadan float64
	var tinggiBadan float64
	var hasil float64

	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)

	hasil = (beratBadan / (tinggiBadan * tinggiBadan)) 

	fmt.Printf("Hasil BMI anda adalah: %.2f", hasil)
}