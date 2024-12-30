package main

import "fmt"

func main() {

	var totalBelanja float64
	var asesmen bool

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Apakah anda mengikuti asesmen CLO?(true/false): ")
	fmt.Scan(&asesmen)

	if asesmen {
		diskon := totalBelanja * 35/100
		fmt.Println(totalBelanja-diskon)
	} else{
		fmt.Print(totalBelanja)
	} 

}