package main

import "fmt"

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func printPessoa(pessoa Pessoa) {

	fmt.Printf("Nome: %s\nIdade: %d\nEndereço: ", pessoa.nome, pessoa.idade)
	fmt.Print(pessoa.endereco)

}

func main() {

	e := Endereco{"SQN", 203, "Brasília", "DF"}
	p := Pessoa{"Gabriel", 17, e}

	printPessoa(p)

}
