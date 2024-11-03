package main

import "fmt"

func main() {

	var nilai int
	var predikat string

	fmt.Print("Masukkan Nilai Ujian: ")
	fmt.Scan(&nilai)

	if nilai > 90 {
		predikat = "A"
	} else if nilai >= 80 && nilai <=90 {
		predikat = "AB"
	} else if nilai >= 70 && nilai <=80 {
		predikat = "B"
	} else if nilai < 70 {
		predikat = "C"
	}

	fmt.Println("Predikat nilai ujian anda adalah: ", predikat)
}