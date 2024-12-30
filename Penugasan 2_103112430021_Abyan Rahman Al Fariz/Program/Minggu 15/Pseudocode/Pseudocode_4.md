Program Hitung Belanja Setelah Diskon

Kamus
    x, y, z : integer
    belanja, diskon : float32

Algoritma
    Input x, y, z

    diskon = (float32(z) * float32(x)) / 100

    If diskon > float32(y)
        diskon = float32(y)
    End If

    belanja = float32(z) - diskon

    Output belanja in Rupiah with 0 decimal places
End Program
