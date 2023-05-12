package main

import "fmt"

// Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas". O campo "notas" deve ser um slice de float64,
// representando as notas que o aluno tirou em uma determinada disciplina. Escreva funções que permitam adicionar ou
// remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func adicionarNotas(conjNotas []float64, novaNota float64) []float64 {

	conjNotas = append(conjNotas, novaNota)
	return conjNotas

}

func removerNotas(conjNotas []float64, notaIndesejada float64) []float64 {

	for i := 0; i < len(conjNotas); i++ {

		if conjNotas[i] == notaIndesejada {

			conjNotas = append(conjNotas[:i], conjNotas[i+1:]...)

		}

	}

	return conjNotas

}

func media(conjNotas []float64) float64 {

	soma := 0.0

	for _, ranConjNotas := range conjNotas {

		soma += ranConjNotas

	}

	return soma / float64(len(conjNotas))

}

func printInfo(aluno Aluno) {

	fmt.Println("O nome do seu aluno é:", aluno.Nome)
	fmt.Println("Sua idade é:", aluno.Idade)
	fmt.Println("Sua média é:", media(aluno.Notas))

}

func main() {

	G := Aluno{

		Nome:  "Gabriel",
		Idade: 17,
		Notas: []float64{7, 6, 10, 9, 8},
	}

	G.Notas = adicionarNotas(G.Notas, 5)
	fmt.Println("Agora suas notas são:", G.Notas)

	fmt.Printf("\n")

	G.Notas = removerNotas(G.Notas, 6)
	fmt.Println("Agora suas notas são:", G.Notas)

	fmt.Printf("\n")

	mediaNotas := media(G.Notas)
	fmt.Println("Sua média é:", mediaNotas)

	fmt.Printf("\n")

	printInfo(G)

}