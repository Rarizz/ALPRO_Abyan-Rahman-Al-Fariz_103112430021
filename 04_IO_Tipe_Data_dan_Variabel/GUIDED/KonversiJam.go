package main

import "fmt"

func main () {
	var waktu int
	var jam int
	var menit int
	var detik int

	fmt.Print("Masukkan waktu dalam Detik: ")
	fmt.Scan(&waktu)

	jam = waktu / 3600
	menit = (waktu % 3600) / 60
	detik = waktu % 60

	fmt.Println(jam, "jam", menit, "menit", detik, "detik")
}