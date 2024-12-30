package main

import "fmt"

func main() {

	var a,b,c,d int

	fmt.Print("Masukkan jumlah gol tim pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah gol tim kedua: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan jumlah gol tim ketiga: ")
	fmt.Scan(&c)
	fmt.Print("Masukkan jumlah gol tim keempat: ")
	fmt.Scan(&d)

	max := a
	min := a

	if b > max {
		max = b
	} else {
		min = b
	}

	if c > max {
		max = c
	} else if c < min {
		min = c
	}

	if d > max {
		max = d
	} else if d < min {
		min = d
	}

	fmt.Println("Gol terbanyak: ", max)
	fmt.Println("Gol tersedikit: ", min)
}