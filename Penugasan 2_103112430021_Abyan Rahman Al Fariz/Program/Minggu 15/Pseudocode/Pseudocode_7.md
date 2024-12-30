Program Menentukan Modus

Kamus
    x, j, bilangan, nX, nZero : integer

Algoritma
    Input x

    nX = 0
    nZero = 0

    For j = 1 to 9
        Input bilangan

        If bilangan == x
            nX = nX + 1
        Else
            nZero = nZero + 1
        End If
    End For

    If nX > nZero
        Output "Modus = " and x
    Else
        Output "Modus = " and nZero
    End If
End Program
