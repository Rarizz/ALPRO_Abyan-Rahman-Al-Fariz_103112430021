Program Cek Saldo Akhir

Kamus
    saldo, transaksi : integer

Algoritma
    Input saldo

    Repeat
        Input transaksi

        If transaksi == 0
            Break
        End If

        saldo = saldo + transaksi
    Until False

    Output "Saldo akhir: ", saldo
End Program
