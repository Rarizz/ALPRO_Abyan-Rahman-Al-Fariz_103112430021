Program Hitung Diskon dan Cashback

Kamus
    totalBelanja : integer
    kartu, dapatKartu, diskon, cashback : boolean
    totalBayar : integer

Algoritma
    Input totalBelanja
    Input kartu

    totalBayar = totalBelanja

    If kartu
        dapatKartu = true
    Else
        dapatKartu = false
    End If

    If totalBelanja >= 200000 and dapatKartu
        diskon = true
        cashback = true
    Else If totalBelanja >= 100000
        diskon = true
        cashback = false
    Else
        diskon = false
        cashback = false
    End If

    If diskon
        totalBayar = totalBayar - (totalBayar * 10 / 100)
    End If

    If cashback
        totalBayar = totalBayar - 75000
    End If

    Output "Kartu?", dapatKartu
    Output "Diskon?", diskon
    Output "Cashback?", cashback
    Output "Rp.", totalBayar
End Program
