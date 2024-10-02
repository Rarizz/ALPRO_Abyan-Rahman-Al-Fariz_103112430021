package main

import "fmt"


func main() {
 var (
satu, dua, tiga string
temp string
 )
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&satu)
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&dua)
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&tiga)
 fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
 temp = satu
 satu = dua
 dua = tiga
 tiga = temp
 fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

//Codingan diatas mengambil 3 input string sesuai urutan dari user, 
//Kemudian variable tersebut ditukar urutannya menggunakan variable temp.
//Setelah ditukar, maka program tersebut akan menampilkan urutan yang baru.