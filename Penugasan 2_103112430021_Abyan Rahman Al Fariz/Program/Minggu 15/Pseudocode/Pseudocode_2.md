Program Cek Waktu Selesai

Kamus
    datang, jdatang, mdatang, durasi, jselesai, mselesai : integer

Algoritma
    Input jdatang, mdatang, durasi

    datang = jdatang * 60 + mdatang + durasi
    jselesai = datang / 60
    mselesai = datang % 60

    If jselesai < 20 or (jselesai == 20 and mselesai == 0)
        Output jselesai and mselesai
    Else
        Output "Silahkan datang kembali besok"
    End If
End Program
