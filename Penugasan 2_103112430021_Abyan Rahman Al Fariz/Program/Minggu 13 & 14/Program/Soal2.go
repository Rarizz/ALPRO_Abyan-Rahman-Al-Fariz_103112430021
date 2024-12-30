    package main

    import "fmt"

    func main() {
        var desimal int
        var biner string

        fmt.Scan(&desimal)

        if desimal == 0 {
            fmt.Println("0")
            return
        }

        for desimal > 0 {
            if desimal%2 == 0 {
                biner = "0" + biner
            } else {
                biner = "1" + biner
            }
            desimal /= 2
        }

        fmt.Println(biner)
    }
