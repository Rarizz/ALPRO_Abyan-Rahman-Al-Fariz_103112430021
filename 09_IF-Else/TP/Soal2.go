package main

import "fmt"

func main() {

	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	if len(huruf) != 1 {
		fmt.Println("Masukkan satu huruf saja.")
		return
	}
	
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Println("Huruf vokal")
	} else {
		fmt.Println("Huruf konsonan")
	}
}