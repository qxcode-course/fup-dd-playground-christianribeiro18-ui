package main
import "fmt"
func main() {
    var A, B int
    var c int
    fmt.Scan(&A, &B)
   
    c= A-B
    if c < 0 {
        c*=-1
    }
    fmt.Println(c)
    }
