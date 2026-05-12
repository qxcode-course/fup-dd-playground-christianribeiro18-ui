package main

import (
	"fmt"
	"math"
)

func aplicarCeil(valor float64) int {
	return int(math.Ceil(valor))
}

func aplicarFloor(valor float64) int {
	return int(math.Floor(valor))
}
func aplicarRound(valor float64) int {
	return int(math.Round(valor))
}

func main() {
	var operacao string
	var valor float64

	fmt.Scan(&operacao)
	fmt.Scan(&valor)

	switch operacao {
	case "c":
		fmt.Println(aplicarCeil(valor))
	case "f":
		fmt.Println(aplicarFloor(valor))
	case "r":
		fmt.Println(aplicarRound(valor))
	}
}