package main
import "fmt"
func main() {
    var A, B, C int
    fmt.Scan(&A, &B, &C)
    if A == B && A== C{
        fmt.Println(3)
   } else if A== B || B== C || C== A{
    fmt.Println(2)
   } else  {
    fmt.Println(0)
   }
}