package main

import "fmt"

func main() {

	var n, m, x, y int

	fmt.Print("Masukkan n, m, x, y: ")
	fmt.Scan(&n, &m, &x, &y)

	cangkir := 0

	for x <= n && y <= m {
		n -= x
		m -= y
		cangkir++
	}

	fmt.Println(cangkir)
}