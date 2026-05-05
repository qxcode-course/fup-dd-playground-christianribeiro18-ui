package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)
    if A>=B {
        fmt.Println(A)
    } else if B>A {
        fmt.Println(B)
    }
}
