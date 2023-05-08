package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calcularAreaDoCirculo(circulo Circulo) float64 {

	return math.Pi * circulo.raio * circulo.raio

}

func main() {

	circulo := Circulo{42.43}

	fmt.Println("A área do círculo é: ", calcularAreaDoCirculo(circulo))

}
