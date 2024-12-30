package main

import "fmt"

func main() {

	var n, lebar, max int

	fmt.Print("Masukkan bilangan yang ingin di inputkan: ")
	fmt.Scan(&n)

	max = 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&lebar)

		if lebar > max {
			max = lebar
		}
	}

	fmt.Println(max)
}