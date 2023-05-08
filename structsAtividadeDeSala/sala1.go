package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func calcularRetangulo(retangulo Retangulo) {

	fmt.Println("Sua área é: ", retangulo.largura*retangulo.altura)

}

func main() {

	var r Retangulo

	fmt.Println("Informe a largura do seu retangulo.")
	fmt.Scanln(&r.largura)

	fmt.Println("Informe a altura do seu retangulo.")
	fmt.Scanln(&r.altura)

	calcularRetangulo(r)

}
