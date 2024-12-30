Program Cari Digit Terbesar

Kamus
    angka, terbesar, digit : integer

Algoritma
    Input angka

    terbesar = 0

    While angka > 0
        digit = angka % 10

        If digit > terbesar
            terbesar = digit
        End If

        angka = angka / 10
    End While

    Output terbesar
End Program
