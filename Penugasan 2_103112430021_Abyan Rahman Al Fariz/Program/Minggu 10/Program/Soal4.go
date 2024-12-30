package main

import "fmt"

func main() {

	var h1, m1, h2, m2 int

	fmt.Print("Masukkan waktu mulai parkir(jam:menit): ")
	fmt.Scan(&h1, &m1)
	fmt.Print("Masukkan waktu selesai parkir(jam:menit): ")
	fmt.Scan(&h2, &m2)

	jamMulai := h1*60 + m1
	jamSelesai := h2*60 + m2

	if jamSelesai < jamMulai {
		jamSelesai += 12 * 60 
	}

	durasi := jamSelesai - jamMulai

	jam := durasi / 60
	menit := durasi % 60

	fmt.Print("Durasi parkir: ", jam, " jam ", menit, " menit")
}