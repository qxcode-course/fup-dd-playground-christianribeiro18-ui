package main
import "fmt"
func main() {
    var A int
    fmt.Scan(&A)
    if A % 7 ==0 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
