package main

import "fmt"

func main() {

	var PanjangSisi int
	var volume int

	fmt.Print("Masukkan Panjang sisi kubus: ")
	fmt.Scan(&PanjangSisi)

	volume = PanjangSisi * PanjangSisi * PanjangSisi

	fmt.Println("Volume kubus adalah: ", volume)
}