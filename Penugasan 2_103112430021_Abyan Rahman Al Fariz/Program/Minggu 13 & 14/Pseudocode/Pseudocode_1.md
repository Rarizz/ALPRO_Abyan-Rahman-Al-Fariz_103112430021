Program Cek Bilangan Prima

Kamus
    bilangan, i : integer

Algoritma
    Input bilangan

    If bilangan < 2
        Output "false"
        Return
    End If

    For i = 2 to sqrt(bilangan)
        If bilangan % i == 0
            Output "false"
            Return
        End If
    End For

    Output "true"
End Program
