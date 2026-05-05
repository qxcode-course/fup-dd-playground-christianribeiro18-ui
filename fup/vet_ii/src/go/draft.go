package main
import "fmt"
func main() {
    var numero int
    fmt.Scan(&numero)
    var arr []int = make([]int, numero)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    for i := range arr {
        fmt.Println(arr[i])
    }
}
