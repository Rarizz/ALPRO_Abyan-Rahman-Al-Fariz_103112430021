Program Menghitung Jumlah Mobil

Kamus
    n, jumlahMobil : integer
    kapasitasMobil : integer

Algoritma
    Input n
    kapasitasMobil = 7
    jumlahMobil = 0

    Jika n lebih besar dari 0 dan n kurang dari atau sama dengan kapasitasMobil
        jumlahMobil = 1
    Jika tidak, jika n lebih besar dari kapasitasMobil
        jumlahMobil = n dibagi kapasitasMobil
        Jika n modulo kapasitasMobil tidak sama dengan 0
            jumlahMobil = jumlahMobil + 1
    Akhir jika

    Output jumlahMobil
End Program