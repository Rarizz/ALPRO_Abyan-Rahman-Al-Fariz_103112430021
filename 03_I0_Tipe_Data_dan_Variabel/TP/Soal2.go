package main

import "fmt"

func main() {

	var angkapertama, angkakedua float64
	var operator string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angkapertama)

	fmt.Print("Masukkan Operator yang ingin dimasukkan (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angkakedua)

	switch operator {

		case "+":
			fmt.Printf("Hasil: %.2f\n", angkapertama+angkakedua)

		case "-":
			fmt.Printf("Hasil: %.2f\n", angkapertama-angkakedua)	

		case "*":				
			fmt.Printf("Hasil: %.2f\n", angkapertama*angkakedua)	

		case "/":
			if angkakedua == 0 {
			fmt.Printf("Hasil: %.2f\n", angkapertama/angkakedua)
			} else {
				fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan")
			}	

		default:	
			fmt.Println("Pakai operator dulu bro!")	

	}
}