package main

import "fmt"

func main() {

	var x,y int

	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	hasilBagi := 0

	for x >= y {
		x -= y
        hasilBagi++
	}

	fmt.Println(hasilBagi) 
}