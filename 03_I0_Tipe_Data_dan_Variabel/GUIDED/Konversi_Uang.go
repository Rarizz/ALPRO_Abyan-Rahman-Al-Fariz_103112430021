package main

import "fmt"

func main() {
	
	var rupiah int
	var konversi int

	fmt.Print("Masukkan jumlah rupiah: ")
	fmt.Scan(&rupiah)
	
	konversi = rupiah / 15000
	fmt.Println("Hasil konveris dollar:", konversi, "dollar")
}