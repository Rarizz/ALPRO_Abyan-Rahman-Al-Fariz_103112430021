Program Konversi Jam 24 ke Jam 12

Kamus
    jam12, jam24 : integer

Algoritma
    Input jam24

    If jam24 % 12 == 0
        jam12 = 12
    Else
        jam12 = jam24 % 12
    End If

    If jam24 < 12
        Output jam12 and " AM"
    Else
        Output jam12 and " PM"
    End If
End Program
