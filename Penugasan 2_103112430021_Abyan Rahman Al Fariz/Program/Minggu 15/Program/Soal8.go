package main

import "fmt"

func main() {

	var x, minggu int
	
	fmt.Scan(&x)

	minggu = x / 7

	if x%7 == 0 {
		fmt.Println("Minggu ke-", minggu)
	} else {
		fmt.Println("Minggu ke-", minggu+1)
	}
}