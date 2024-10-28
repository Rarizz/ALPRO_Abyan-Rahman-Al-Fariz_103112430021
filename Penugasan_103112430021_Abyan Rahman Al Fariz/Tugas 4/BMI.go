package main

import "fmt"

func main(){
	var nama string
	var beratBadan float64
	var tinggiBadan float64
	var hasil float64

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)

	hasil = (beratBadan / (tinggiBadan * tinggiBadan)) 

	fmt.Println("Informasi BMI")
	fmt.Println("Nama: ", nama)
	fmt.Printf("Berat badan: %.2f kg\n", beratBadan)
	fmt.Printf("Tinggi badan: %.2f m\n", tinggiBadan)
	fmt.Printf("BMI: %.2f", hasil)
}