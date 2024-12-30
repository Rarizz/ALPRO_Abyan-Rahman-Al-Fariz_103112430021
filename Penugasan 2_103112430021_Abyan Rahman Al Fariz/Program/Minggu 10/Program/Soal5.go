package main

import "fmt"

func main() {

	var totalBelanja int
	var kartu, dapatKartu, diskon, cashback bool

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Apakah bersedia dibuatkan kartu?(true/false): ")
	fmt.Scan(&kartu)

	totalBayar := totalBelanja

	if kartu {
        dapatKartu = true
    } else {
		dapatKartu = false
	}

	if totalBelanja >= 200000 && dapatKartu { 
		diskon = true
		cashback = true
	} else if totalBelanja >= 100000 {
		diskon = true
		cashback = false
	} else {
		diskon = false
		cashback = false
	}

	if diskon {
		totalBayar -= totalBayar * 10 / 100
	}

	if cashback {
		totalBayar -= 75000
	}

	fmt.Println("Kartu?",dapatKartu)
	fmt.Println("Diskon?",diskon)
	fmt.Println("Cashback?",cashback)
	fmt.Println("Rp.",totalBayar)
}