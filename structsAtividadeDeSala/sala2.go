package main

import "fmt"

type Livro struct {
	titulo          string
	autor           string
	anoDePublicacao int
	fa              string
}

func printLivro(livro Livro) {

	fmt.Printf("Título: %s\nAutor: %s\nAno de publicação: %d\nFã: %s", livro.titulo, livro.autor, livro.anoDePublicacao, livro.fa)

}

func main() {

	livro := Livro{"Cem Anos de Solidão", "Gabriel García Márquez", 1982, "Isabelle"}

	printLivro(livro)

}
