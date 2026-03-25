package main

import (
	"fmt"
	"math"
)

func main() {
	var valor1, valor2, valor3 float32
	fmt.Scanf("%float32", &valor1)
	fmt.Scanf("%float32", &valor2)
	fmt.Scanf("%float32", &valor3)
	p := (valor1 + valor2 + valor3) / 2
	area := (p * (p - valor1) * (p - valor2) * (p - valor3))
	t := math.Sqrt(float64(area))
	fmt.Printf("%.2f\n", t)

}
