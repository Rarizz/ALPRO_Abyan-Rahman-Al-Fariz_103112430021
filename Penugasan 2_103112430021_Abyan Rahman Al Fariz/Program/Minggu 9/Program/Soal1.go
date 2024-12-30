package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	kapasitasMobil := 7
	jumlahMobil := 0

	if n > 0 && n<= kapasitasMobil {
		jumlahMobil = 1
	} else if n > kapasitasMobil {
		jumlahMobil = n / kapasitasMobil
		if n%kapasitasMobil != 0 {
			jumlahMobil++
		}
	}

	fmt.Println("Jumlah mobil yang dibutuhkan: ", jumlahMobil)

}