Program Cek Gol Terbanyak dan Tersedikit

Kamus
    a, b, c, d : integer
    max, min : integer

Algoritma
    Input a
    Input b
    Input c
    Input d

    max = a
    min = a

    If b > max
        max = b
    Else
        min = b
    End If

    If c > max
        max = c
    Else If c < min
        min = c
    End If

    If d > max
        max = d
    Else If d < min
        min = d
    End If

    Output "Gol terbanyak: ", max
    Output "Gol tersedikit: ", min
End Program
