Program Cek Pengisian Tangki

Kamus
    t, v, total : integer
    selesai : boolean

Algoritma
    Input t

    total = 0
    selesai = false

    Repeat
        Input v

        total = total + v

        selesai = total >= t
        Output selesai
    Until selesai = true

End Program
