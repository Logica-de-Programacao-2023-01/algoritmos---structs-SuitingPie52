package main

import (
	"bufio"
	"fmt"
	"os"
)

// Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som".
// Escreva funções que permitam modificar o som que o animal faz
// e uma função que imprima as informações do animal e o som que ele faz.

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func mudarSom(animal Animal) Animal {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Qual o novo som do seu animal?")
	scanner.Scan()
	novoSom := scanner.Text()

	animal.Som = novoSom

	fmt.Println("Pronto! Agora o som desse animal é: ", animal.Som)
	return animal

}

func infoAnimal(animal Animal) {

	fmt.Println("Seu nome é: ", animal.Nome)
	fmt.Println("Sua espécie é: ", animal.Especie)
	fmt.Println("Sua idade é: ", animal.Idade)
	fmt.Println("E seu som é: ", animal.Som)

}

func main() {

	dog := Animal{

		Nome:    "Tutti",
		Especie: "Cachorro (Shi-Tzu)",
		Idade:   1,
		Som:     "Au Au",
	}

	cat := Animal{
		Nome:    "Frajolinha",
		Especie: "Frajola",
		Idade:   0,
		Som:     "Miauuuuuuuuuuuuuuuu",
	}

	infoAnimal(dog)

	fmt.Printf("\n\n")

	cat = mudarSom(cat)

	fmt.Printf("\n\n\n\n%s", cat.Som)
}
