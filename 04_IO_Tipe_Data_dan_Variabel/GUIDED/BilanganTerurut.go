package main

import "fmt"

func main() {

	var bilangan int
	var digitsatu int
	var digitandua int
	var digitketiga int

	fmt.Print("Masukkan bilangan 3 digit: ")
	fmt.Scan(&bilangan)

	digitsatu = bilangan / 100
	digitandua = (bilangan % 100) / 10
	digitketiga = bilangan % 10
	
	fmt.Println(digitsatu <= digitandua && digitandua <= digitketiga)
}