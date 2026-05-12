package main
import "fmt"
func main() {
    var nota1, nota2, notaFinal int
    fmt.Scan(&nota1,&nota2, &notaFinal)
    if (nota1+nota2)/2 >=7 {
        fmt.Println("aprovado")
    } else if (nota1+nota2)/2 <4 {
        fmt.Println("reprovado")
    } else if (notaFinal+((nota1+nota2)/2))/2 >=5 {
        fmt.Println("aprovado na final")
    } else {
        fmt.Println("reprovado na final")
    }
}
