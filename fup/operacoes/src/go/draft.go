package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)
    fmt.Println(A+B)
    fmt.Println(A-B)
    fmt.Println(A*B)
    fmt.Printf("%.2f\n", float64 (A)/float64(B))
    fmt.Println(A%B)
    
}
