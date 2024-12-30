Program Konversi Desimal ke Biner

Kamus
    desimal : integer
    biner : string

Algoritma
    Input desimal

    If desimal == 0
        Output "0"
        Return
    End If

    biner = ""

    While desimal > 0
        If desimal % 2 == 0
            biner = "0" + biner
        Else
            biner = "1" + biner
        End If

        desimal = desimal / 2
    End While

    Output biner
End Program
