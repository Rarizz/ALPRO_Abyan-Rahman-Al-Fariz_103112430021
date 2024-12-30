package main

import "fmt"

func main() {

	var a,b,c int

	fmt.Print("Masukkan sisi pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan sisi kedua: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan sisi ketiga: ")
	fmt.Scan(&c)

	if a == b && b == c{
		fmt.Println("Segitiga sama sisi")
	} else if a == b || b == c || a == c {
        fmt.Println("Segitiga sama kaki")
    } else {
        fmt.Println("Segitiga sembarang")
	}
}