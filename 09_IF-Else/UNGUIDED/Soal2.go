// 1. Keluaran dari program adalah D. Ini karena variabel nam diubah beberapa kali oleh logika dalam program dan akhirnya hanya mencocokkan kondisi terakhir.
// 2. Kesalahan dari program tersebut adalah:
//    a. Penggunaan variabel yang belum sesuai, variabel NAM yang digunakan untuk menyimpan nilai huruf (NMK), yang seharusnya menggunakan variabel nmk.
//    b. Penggunaan struktur if dan percabangan yang kurang tepat.
//    c. Penggunaan variabel nam digunakan untuk menyimpan nilai huruf, yang semestinya diletakkan pada variabel string (nmk).

// 3. Berikut adalah perbaikan codingan nya:
package main

import (
    "fmt"
)

func main() {
	var nam float64
    var nmk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah:", nmk)
}
