package main
import "fmt"
func main() {
    var a,b,c int
    fmt.Scan(&a, &b, &c)
    if a>b && b>c {
        fmt.Println(b)
    } else if b>c && c>a {
        fmt.Println(c)
    } else if c>a && a>b {
        fmt.Println(a)
    } else if c >b && c<a {
        fmt.Println(c)
    } else if b>a && b<c {
        fmt.Println(b)
    } else if a<b && a>c {
        fmt.Println(a)
    } 
}
