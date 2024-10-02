package main

import "fmt"

func main() {
	var (
		nama string
		nim int
		kelas string	
	)

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Kelas Anda: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukkan NIM Anda: ")
	fmt.Scanln(&nim)

	fmt.Println("Perkenalkan saya adalah", nama,"salah satu mahasiswa Prodi S1-", kelas,"dengan NIM", nim)
}