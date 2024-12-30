Program Cek Bilangan Konsekutif

Kamus
    x, bilangan, selisih : integer
    bilanganKonsekutif : boolean

Algoritma
    Input x

    bilanganKonsekutif = true
    selisih = -1

    Repeat
        bilangan = x % 10

        If selisih != -1 and (selisih - bilangan != 1 and bilangan - selisih != 1)
            bilanganKonsekutif = false
            Break
        End If

        selisih = bilangan
        x = x / 10
    Until x <= 0

    Output bilanganKonsekutif
End Program
