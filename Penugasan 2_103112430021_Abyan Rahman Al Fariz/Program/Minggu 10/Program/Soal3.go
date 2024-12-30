package main

import "fmt"

func main() {

	var bulanIni, bulanSebelumnya float64

	fmt.Print("Masukkan keuntungan bulan sebelumnya: ")
	fmt.Scan(&bulanSebelumnya)
	fmt.Print("Masukkan keuntungan bulan ini: ")
	fmt.Scan(&bulanIni)

	if bulanIni > bulanSebelumnya {
		selisih := bulanIni - bulanSebelumnya
		fmt.Println("Selisih keuntungan bulan ini dan bulan sebelumnya: ", selisih)
	} else if bulanIni < bulanSebelumnya {
        selisih := bulanSebelumnya - bulanIni
        fmt.Println("Selisih keuntungan bulan ini dan bulan sebelumnya: ", selisih)
    }
}