Program Hitung Tipe Masukan

Kamus
    masukan : string
    jumlahA, jumlahB, jumlahC : integer

Algoritma
    jumlahA = 0
    jumlahB = 0
    jumlahC = 0

    Repeat
        Input masukan

        If masukan == "A"
            jumlahA = jumlahA + 1
        Else If masukan == "B"
            jumlahB = jumlahB + 1
        Else If masukan == "C"
            jumlahC = jumlahC + 1
        Else
            Break
        End If
    Until False

    Output "Tipe A:", jumlahA
    Output "Tipe B:", jumlahB
    Output "Tipe C:", jumlahC
End Program
