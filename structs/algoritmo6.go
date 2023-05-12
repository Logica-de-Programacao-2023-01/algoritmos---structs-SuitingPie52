package main

import (
	"bufio"
	"fmt"
	"os"
)

// Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade".
// Escreva funções que permitam aumentar ou diminuir o salário do funcionário em uma determinada porcentagem
// e uma função que calcule o tempo de serviço do funcionário
//(considerando que ele começou a trabalhar aos 18 anos).

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func menuFuncionarios() int {

	x := 0

	fmt.Printf("Escolha o que você quer fazer:\n\n1- Alterar salário\n2- Ver tempo de trabalho.\nOutro - Sair do algoritmo\n\n")
	fmt.Scanln(&x)
	return x

}

func alterarSalario(lessOrMore int, porcentagem float64, funcionario Funcionario) (float64, error) {

	for porcentagem < 0 {

		fmt.Println("Sua porcentagem é negativa. Digite outra porcentagem:")
		fmt.Scanln(&porcentagem)

	}

	for {

		if lessOrMore == 1 {

			funcionario.Salario += funcionario.Salario * (porcentagem / 100)
			return funcionario.Salario, nil

		} else if lessOrMore == 2 {

			funcionario.Salario -= funcionario.Salario * (porcentagem / 100)
			return funcionario.Salario, nil

		} else {

			fmt.Println("Você digitou um comando que não existe. Digite o que você quer fazer novamente:")
			fmt.Scanln(&lessOrMore)

		}
	}
}

func tempoDeTrabalho(nome string, conjFuncionarios []Funcionario) {

	anosDeTrabalho := 0

	for _, ranConjFuncionarios := range conjFuncionarios {

		if ranConjFuncionarios.Nome == nome {

			anosDeTrabalho = ranConjFuncionarios.Idade - 18

			if anosDeTrabalho == 0 {

				fmt.Println("O funcionário começou a trabalhar esse ano!")
				return

			} else if anosDeTrabalho != 0 {

				fmt.Println("O tempo de trabalho desse funcionário é: ", anosDeTrabalho, "anos!")
				return

			}

		}

	}

	fmt.Println("Não existe esse funcionário.")

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	BELLE := Funcionario{
		Nome:    "Isabelle",
		Salario: 1200,
		Idade:   20,
	}

	JHONY := Funcionario{
		Nome:    "João Eduardo",
		Salario: 3000,
		Idade:   18,
	}

	PAI := Funcionario{
		Nome:    "Fábio",
		Salario: 10000,
		Idade:   60,
	}

	conjFuncionarios := []Funcionario{BELLE, JHONY, PAI}
	nomeFuncionario := ""

	fmt.Printf("Este são os seus funcionários:\n\n")

	for _, ranFuncionarios := range conjFuncionarios {

		fmt.Println(ranFuncionarios.Nome)

	}

	fmt.Printf("\n-----------------------------------------\n\n")

	switch menuFuncionarios() {

	case 1:

		comando := 0
		fmt.Printf("Você deseja aumentar ou diminuir o salário?\n\n1- Aumentar.\n2- Diminuir.\n\n")
		fmt.Scanln(&comando)

		porcentagem := 0.0
		fmt.Println("Em qual porcentagem?")
		fmt.Scanln(&porcentagem)

		fmt.Println("De qual funcionário?")
		scanner.Scan()
		nomeFuncionario = scanner.Text()

		for i := 0; i < len(conjFuncionarios); i++ {

			if nomeFuncionario == conjFuncionarios[i].Nome {

				conjFuncionarios[i].Salario, _ = alterarSalario(comando, porcentagem, conjFuncionarios[i])
				fmt.Println("Agora o salário desse funcionário é: ", conjFuncionarios[i].Salario, nil)
				break

			}

		}

		fmt.Printf("\n")

	case 2:

		fmt.Println("De qual funcionário:")
		scanner.Scan()
		nomeFuncionario = scanner.Text()

		tempoDeTrabalho(nomeFuncionario, conjFuncionarios)

	default:

		break

	}

}
