Program Cek Jumlah Digit

Kamus
    x, jumlah, digit : integer

Algoritma
    Input x

    jumlah = 0

    Repeat
        digit = x % 10
        Output digit
        jumlah = jumlah + digit
        x = x / 10
    Until x <= 0

    Output "Jumlah digit: ", jumlah
End Program
