package main

import "fmt"

func main(){

	var celsius float32

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := (celsius * 4 / 5)
	kelvin := celsius + 273.15

	fmt.Println("Suhu Celsius dalam derajat Fahrenheit adalah" , fahrenheit)
	fmt.Println("Suhu Celsius dalam derajat Reamur adalah" , reamur)
	fmt.Println("Suhu Celsius dalam derajat Kelvin adalah" , kelvin)
}