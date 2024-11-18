package main

import "fmt"

func main() {

	var beratGram int


	fmt.Print("Masukkan berat parsel dalam satuan gram: ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaPerKg := 10000
	totalBiaya := kg * biayaPerKg
	sisaBiaya := 0

	if kg > 10 {
		sisaBiaya = 0
	} else {
		if sisaGram >= 500 {
			sisaBiaya = sisaGram * 5
        } else if sisaGram > 500{
			sisaBiaya = sisaGram * 15
	    }
	}

	totalBiaya = totalBiaya + sisaBiaya

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*biayaPerKg, sisaBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	
}