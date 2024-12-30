package main

import "fmt"

func main() {

	var t1,t2,t3,t4,t5 float64

	fmt.Print("Masukkan nilai t1: ")
	fmt.Scan(&t1)
	fmt.Print("Masukkan nilai t2: ")
	fmt.Scan(&t2)
	fmt.Print("Masukkan nilai t3: ")
	fmt.Scan(&t3)
	fmt.Print("Masukkan nilai t4: ")
	fmt.Scan(&t4)
	fmt.Print("Masukkan nilai t5: ")
	fmt.Scan(&t5)

	if t1 < t2 && t2 < t3 && t3 < t4 && t4 < t5 {
		fmt.Println("Stabil naik")
	} else if t1 > t2 && t2 > t3 && t3 > t4 && t4 > t5 {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}