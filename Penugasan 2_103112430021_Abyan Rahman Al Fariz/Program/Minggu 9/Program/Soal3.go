package main

import (
    "fmt"
)

func main() {

	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" || huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E"{
		fmt.Println("Huruf bukan konsonan")
	} else if huruf >= "A" && huruf <= "Z" || huruf >= "a" && huruf <= "z"{
		fmt.Println("Huruf konsonan")
	} else {
		fmt.Println("Bukan huruf")
	}
}