package main

import "fmt"

func isLeapYear(year int) bool {
	// Memeriksa apakah tahun kabisat
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func main() {
	var tahun int
	
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	fmt.Println(isLeapYear(tahun))
}