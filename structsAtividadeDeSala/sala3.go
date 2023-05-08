package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func calcularMedia(aluno Aluno) float64 {

	soma := 0.0

	for _, ranNotas := range aluno.notas {

		soma += ranNotas

	}

	return soma / float64(len(aluno.notas))

}

func main() {

	nota := 0.0
	notas := []float64{}

	for {

		fmt.Println("Informe uma nota: (Digite -1 para sair) ")
		fmt.Scanln(&nota)

		if nota == -1 {
			break
		}

		notas = append(notas, nota)
	}

	aluno := Aluno{"Gabriel", 17, notas}
	media := calcularMedia(aluno)

	fmt.Println("Sua média é: ", media)

}
