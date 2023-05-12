package main

import "fmt"

// Crie uma struct chamada Viagem com os campos
// "origem", "destino", "data" e "preço".
// Escreva uma função que receba um
// slice de Viagens como parâmetro
// e retorne a viagem mais cara.

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func mostExpensive(conjViagens []Viagem) Viagem {

	maior := conjViagens[0].Preco

	for _, ranConjViagens := range conjViagens {

		if ranConjViagens.Preco > maior {

			maior = ranConjViagens.Preco

		}
	}

	for _, ranConjViagens := range conjViagens {

		if ranConjViagens.Preco == maior {

			return ranConjViagens

		}
	}

	return conjViagens[0]

}

func main() {

	DFtoG := Viagem{

		Origem:  "Brasília",
		Destino: "Guarapari",
		Data:    "17/07",
		Preco:   2130,
	}

	GtoDF := Viagem{

		Origem:  "Guarapari",
		Destino: "Brasília",
		Data:    "01/12",
		Preco:   1210,
	}

	SPtoDF := Viagem{

		Origem:  "São Paulo",
		Destino: "Brasília",
		Data:    "27/7",
		Preco:   4500,
	}

	DFtoSP := Viagem{

		Origem:  "Brasília",
		Destino: "São Paul",
		Data:    "05/05",
		Preco:   3400,
	}

	conjViagens := []Viagem{DFtoG, GtoDF, SPtoDF, DFtoSP}

	fmt.Println("Sua viagem mais cara é:", mostExpensive(conjViagens).Origem, "para", mostExpensive(conjViagens).Destino)

}
