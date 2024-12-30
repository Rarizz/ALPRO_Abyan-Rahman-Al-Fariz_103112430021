Program Cek Digit dalam Angka

Kamus
    x, n, digit : integer
    kebenaran : boolean

Algoritma
    Input x, n

    kebenaran = false

    While n > 0
        digit = n % 10

        If digit == x
            kebenaran = true
            Break
        End If

        n = n / 10
    End While

    Output kebenaran
End Program
