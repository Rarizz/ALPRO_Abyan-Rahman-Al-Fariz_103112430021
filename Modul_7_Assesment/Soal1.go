package main

import "fmt"

func main() {

	var x int
	var y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil:= 0

	for i:=x; i<=y; i++ {

		hasil += i

		fmt.Println("Hasil penjumlahan: ", hasil)
	}
}