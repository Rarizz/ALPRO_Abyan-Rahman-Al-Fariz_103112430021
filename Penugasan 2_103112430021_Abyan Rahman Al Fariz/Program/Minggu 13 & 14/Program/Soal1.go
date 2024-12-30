package main

import "fmt"

func main() {
    var bilangan int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&bilangan)

    if bilangan < 2 {
        fmt.Println("false") 
        return
    }

    for i := 2; i*i <= bilangan; i++ {
        if bilangan%i == 0 { 
            fmt.Println("false") 
            return
        }
    }

    fmt.Println("true") 
}
