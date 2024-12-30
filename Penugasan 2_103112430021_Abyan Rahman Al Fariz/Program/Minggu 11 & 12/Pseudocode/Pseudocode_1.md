Program Cek Login

Kamus
    username, password : string
    usernameBenar, passwordBenar : string
    percobaanGagal : integer

Algoritma
    usernameBenar = "Admin"
    passwordBenar = "Admin"
    percobaanGagal = 0

    Repeat
        Input username
        Input password

        If username == usernameBenar and password == passwordBenar
            Break
        Else
            Output "Username atau Password salah."
            percobaanGagal = percobaanGagal + 1
        End If
    Until False

    Output percobaanGagal, "Percobaan Gagal Login"
End Program
