package main

import (
	"fmt"
)

func main() {

	var a int

	fmt.Print("Masukkan angka yang ingin di faktorialkan: ")
	fmt.Scan(&a)

	faktorial := 1
	for i := 1; i <= a; i++ {
		faktorial *= i
	}
	fmt.Println("Hasil dari faktorial", a, "adalah",faktorial)

}