Program Hitung Tarif Parkir dengan Diskon

Kamus
    tarif, potongan, diskon, tarifAwal : float32
    durasi, kelebihan : integer
    member : string

Algoritma
    Input member, durasi

    If member == "Gold"
        diskon = 0.5
    Else If member == "Silver"
        diskon = 0.25
    Else
        diskon = 0
    End If

    If durasi <= 2
        tarifAwal = durasi * 65000
    Else
        kelebihan = durasi - 2
        tarifAwal = (2 * 65000) + (kelebihan * 20000)
    End If

    potongan = diskon * tarifAwal
    tarif = tarifAwal - potongan

    If member == "None" and durasi <= 2
        Output "IDR" and tarif
    Else If member == "Gold" or member == "Silver"
        Output "IDR" and tarif
    Else
        Output "IDR" and tarif
    End If
End Program
