package main

import "fmt"

func main() {

	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan kewarganegaraan: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "WNI" || kewarganegaraan == "wni" {
		fmt.Println("Selamat anda bisa mengikuti pemilu!")
	} else {
		fmt.Print("Anda tidak bisa mengikuti pemilu karena ")
		if umur < 17 && kewarganegaraan != "WNI" {
			fmt.Println("umur Anda di bawah 17 tahun dan Anda bukan WNI.")
		} else if umur < 17 {
			fmt.Println("umur Anda di bawah 17 tahun.")
		} else if kewarganegaraan != "WNI" {
			fmt.Println("Anda bukan WNI.")
		}
	}
}