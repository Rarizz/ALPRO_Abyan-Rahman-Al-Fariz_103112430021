package main

import "fmt"

func main() {

	var semester int
	var skorEPrT int

	fmt.Print("Masukkan semester: ")
	fmt.Scan(&semester)

	fmt.Print("Masukkan skor EPrT: ")
	fmt.Scan(&skorEPrT)

	if semester <= 8 && skorEPrT >= 500 {
		fmt.Printf("Mahasiswa lulus cumlaude dengan kuliah selama %d semester dan EPrT %d\n", semester, skorEPrT)
	} else {
		fmt.Println("Mahasiswa tidak lulus cumlaude")
	}

}