package main

import "fmt"

func main() {

	var x int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	bilanganKonsekutif := true
	selisih := -1

	for x > 0 {
		bilangan := x%10
		if selisih != -1 && (selisih-bilangan != 1 && bilangan-selisih != 1) {
			bilanganKonsekutif = false
			break
		}
		selisih = bilangan
		x /= 10
	}
	fmt.Println(bilanganKonsekutif)
}