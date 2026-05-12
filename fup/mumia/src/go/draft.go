package main
import "fmt"
func main() {
    var Maria string
    var idade int
    fmt.Scan(&Maria, &idade)
    if idade <12 {
        fmt.Println(Maria, "eh crianca")
    } else if idade <18 {
        fmt.Println(Maria, "eh jovem")
    } else if idade <65 {
        fmt.Println(Maria, "eh adulto")
    } else if idade <1000 {
        fmt.Println(Maria, "eh idoso")
    } else {
        fmt.Println(Maria, "eh mumia")
    }
}
