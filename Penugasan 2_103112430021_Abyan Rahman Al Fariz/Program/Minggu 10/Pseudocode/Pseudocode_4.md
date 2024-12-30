Program Hitung Durasi Parkir

Kamus
    h1, m1, h2, m2 : integer
    jamMulai, jamSelesai, durasi : integer
    jam, menit : integer

Algoritma
    Input h1, m1
    Input h2, m2

    jamMulai = h1 * 60 + m1
    jamSelesai = h2 * 60 + m2

    If jamSelesai < jamMulai
        jamSelesai = jamSelesai + 12 * 60
    End If

    durasi = jamSelesai - jamMulai

    jam = durasi / 60
    menit = durasi % 60

    Output "Durasi parkir: ", jam, " jam ", menit, " menit"
End Program
