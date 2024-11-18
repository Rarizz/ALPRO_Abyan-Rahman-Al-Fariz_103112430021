package main

import (
    "fmt"
)	

func main() {

	var bilangan int

	fmt.Print("Masukkan bilangan 4 digit saja: ")
	fmt.Scan(&bilangan)

	digit1 := bilangan / 1000
	digit2 := (bilangan / 100) %10
	digit3 := (bilangan / 10) %10
	digit4 := bilangan % 10

	if bilangan <= 1000 && bilangan >= 9999 {
		fmt.Println("Bilangan harus 4 digit saja")
		return
	}

	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit pada bilangan", bilangan, "terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4{
		fmt.Println("Digit pada bilangan", bilangan, "terurut mengecil")
	} else {
        fmt.Println("Digit pada bilangan", bilangan, "tidak terurut")
    } 
}