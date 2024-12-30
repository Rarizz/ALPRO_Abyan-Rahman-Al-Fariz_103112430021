Program Cek Pasangan Angka Genap dan Ganjil

Kamus
    x1, x2 : integer
    status : boolean

Algoritma
    Input x1, x2

    status = (x1 % 2 == 0 and x2 % 2 != 0) or (x1 % 2 != 0 and x2 % 2 == 0)

    If status == true
        Output "Berhasil"
    End If
End Program
