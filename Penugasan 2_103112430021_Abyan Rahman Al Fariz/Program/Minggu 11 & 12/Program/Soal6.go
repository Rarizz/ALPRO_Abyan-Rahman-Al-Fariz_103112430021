package main

import "fmt"

func main() {

	var t,v,total int
	var selesai bool

	fmt.Print("Masukkan kapasitas tangki(T): ")
	fmt.Scan(&t)

	for selesai = false; !selesai;{
		fmt.Scan(&v)

		total += v

		selesai = total>= t
		fmt.Println(selesai)
	}
}