package main

import (
	"fmt"
	"math"
)

func main() {

	var(
		jarijari float64
		volume float64
		luasbola float64
	)

		fmt.Print("Masukan jari-jari bola: ")
		fmt.Scan(&jarijari)

		volume = (4.0/3.0) * math.Pi * math.Pow(jarijari, 3)
		luasbola = 4 * math.Pi * math.Pow(jarijari, 2)

		fmt.Printf("Volume bola adalah: %.4f\n", volume)
		fmt.Printf("Luas bola adalah: %.4f\n", luasbola)



}