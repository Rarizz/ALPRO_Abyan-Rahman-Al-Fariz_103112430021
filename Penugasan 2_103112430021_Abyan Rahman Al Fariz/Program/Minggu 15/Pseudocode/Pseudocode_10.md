Program Menentukan Median Bernilai

Kamus
    j, y, bilangan, total : integer

Algoritma
    Input y

    total = 0

    For j = 1 to 9
        Input bilangan
        total = total + bilangan
    End For

    If total >= y * 5
        Output "Median Bernilai " and y
    Else
        Output "Median Bernilai 0"
    End If
End Program
