package main

import "fmt"

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []float64
}

func addRating(rating float64, conjRatings []float64) []float64 {

	conjRatings = append(conjRatings, rating)
	return conjRatings

}

func removeRating(idx int, allRatings []float64) []float64 {

	allRatings = append(allRatings[:idx], allRatings[idx+1:]...)
	return allRatings

}

func averageRating(filme Filme) {

	soma := 0.0

	for _, ranRatings := range filme.avaliacoes {

		soma += ranRatings

	}

	fmt.Println("A nota média do seu filme é: ", soma/float64(len(filme.avaliacoes)))

}

func printInfoFilme(filme Filme) {

	fmt.Println("Título: ", filme.titulo)
	fmt.Println("Diretor: ", filme.diretor)
	fmt.Println("Ano: ", filme.ano)
	fmt.Println("Avaliações: ", filme.avaliacoes)

}

func main() {

	FC := Filme{
		titulo:     "Clube da Luta",
		diretor:    "David Fincher",
		ano:        1999,
		avaliacoes: []float64{8, 7.5},
	}

	FC.avaliacoes = addRating(5, FC.avaliacoes)

	printInfoFilme(FC)

	averageRating(FC)

	fmt.Printf("\n\n")

	SW3 := Filme{
		titulo:     "Star Wars: Episode I - The Phantom Menace",
		diretor:    "George Lucas",
		ano:        1999,
		avaliacoes: []float64{5, 3423423423, 6, 5},
	}

	SW3.avaliacoes = removeRating(1, SW3.avaliacoes)

	printInfoFilme(SW3)

	averageRating(SW3)

}
