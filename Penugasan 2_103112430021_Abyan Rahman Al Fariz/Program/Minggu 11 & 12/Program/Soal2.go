package main

import "fmt"

func main() {

	var saldo, transaksi int

	fmt.Print("Masukkan saldo awal: ")
	fmt.Scan(&saldo)

	for{
		//masukkan 0 jika ingin menyelesaikan transaksi
		fmt.Scan(&transaksi)

		if transaksi == 0 {
            break
        }
		saldo += transaksi
	}

	fmt.Println("Saldo akhir: ", saldo)
}