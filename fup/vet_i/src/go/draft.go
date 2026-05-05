package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    var  arr []int= make([]int, N)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    if N == 0 {
        fmt.Print("\n")
    }

    for i := range arr {
        fmt.Println(arr[i])
    }
}
