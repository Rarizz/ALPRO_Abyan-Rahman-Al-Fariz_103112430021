package main

import "fmt"

func main() {

	var (
		Fahrenheit float32
		Kelvin float32
	)

	fmt.Println("Masukan suhu awal dalam derajat Fahrenheit")
	fmt.Scan(&Fahrenheit)
	
	//rumus konversi dari Fahrenheit ke Kelvin
	Kelvin = (Fahrenheit-32)*5/9 + 273

	fmt.Println("Derajat suhu dalam kelvin adalah", Kelvin)
}