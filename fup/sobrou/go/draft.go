package main
import "fmt"
func main() {
    var p1, p2, p3 float64
    var q1, q2, q3 int
    var dinheiro float64
    fmt.Scan (&q1, &q2, &q3)
    fmt.Scan (&p1, &p2, &p3)
    fmt.Scan (&dinheiro)
    gasto := (float64 (q1)*p1)+ (float64(q2)*p2)+(float64 (q3)*p3)
    troco := float64 (dinheiro - gasto)
    fmt.Printf("%.2f\n", troco)
}
