package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func calcularAreaTriangulo(triangulo Triangulo) float64 {

	return (triangulo.base * triangulo.altura) / 2

}

func main() {

	t := Triangulo{12.5, 23}

	fmt.Println("A área do triângulo é: ", calcularAreaTriangulo(t))
}