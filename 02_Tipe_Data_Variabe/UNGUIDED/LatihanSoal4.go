package main

import "fmt"

func main() {
	var(
		
		konversi int
		f int
	)

	fmt.Println("Masukkan suhu dalam fahrenheit:")
	fmt.Scan(&f)
	
	//rumus buat konversi dari fahrenheit ke celcius
	konversi = (f - 32) * 5 / 9 
	fmt.Println("Suhu dalam celcius adalah:", konversi )
}