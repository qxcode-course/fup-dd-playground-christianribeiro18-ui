package main
import "fmt"
func main() {
    var wifi, login, admin int
    fmt.Scan(&wifi, &login, &admin)
    if !wifi == 1 {
        fmt.Println("you must connect to wifi")
    } else if !login ==1 {
        fmt.Println("you need to login fist")
    } else if !login ==1{
        fmt.Println("you must login as admin")
    }
}