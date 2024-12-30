Program Cek Kondisi untuk Ice Cream

Kamus
    kTidur, tKebun, tMain : integer
    kondisi1, kondisi2, kondisi3 : boolean

Algoritma
    Input kTidur, tMain, tKebun

    kondisi1 = kTidur >= 60 and tMain >= 75 and tKebun >= 60
    kondisi2 = kTidur >= 80 and tKebun >= 80
    kondisi3 = kTidur >= 100

    If kondisi1 or kondisi2 or kondisi3
        Output "Ice cream"
    Else
        Output "Tidak"
    End If
End Program
