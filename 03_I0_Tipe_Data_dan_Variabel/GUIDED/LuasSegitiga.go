package main

import "fmt"

func main (){

	var alas float32
	var tinggi float32
	var luas float32

	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)
	luas = (alas*tinggi)/2
	fmt.Println("Luas segitiga adalah: ", luas)
}