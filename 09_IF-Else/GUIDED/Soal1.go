package main

import "fmt"

func main() {

	var umur int
	var kartuKeluarga bool
	
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan kartu keluarga: ")
	fmt.Scan(&kartuKeluarga)

	if umur >=17 && kartuKeluarga{
		fmt.Print("Anda bisa membuat KTP")
	} else {
		fmt.Print("Belum bisa buat KTP")
	}
}