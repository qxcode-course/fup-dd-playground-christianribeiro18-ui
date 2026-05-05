package main
import "fmt"


func main() {
    var A, B int
    var C string 
    fmt.Scan(&A, &B, &C)
    if C == "+" {
        fmt.Println(A+B)
    } else if C == "-" {
        fmt.Println(A-B)
    } else if C == "*" {
        fmt.Println(A*B)
    } else if C == "/" {
        fmt.Println(A/B)
    }
        

    } 
