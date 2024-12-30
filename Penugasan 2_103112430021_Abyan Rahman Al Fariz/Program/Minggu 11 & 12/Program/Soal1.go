package main

import "fmt"

func main() {

	var username, password string

	usernameBenar := "Admin"
	passwordBenar := "Admin"

	percobaanGagal := 0

	for {
		fmt.Print("Masukkan username: ")
        fmt.Scan(&username)
        fmt.Print("Masukkan password: ")
        fmt.Scan(&password)

		if username == usernameBenar && password == passwordBenar {
			break
		} else {
			fmt.Println("Username atau Password salah.")
			percobaanGagal++
		}
	}
	fmt.Println(percobaanGagal, "Percobaan Gagal Login")
}