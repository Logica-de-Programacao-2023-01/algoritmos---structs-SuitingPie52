package main

import (
	"bufio"
	"fmt"
	"os"
)

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo       string
	Artista      string
	Duracao      int
	Recomendacao string
}

func verifyTitles(musica string, conjPlaylists []Playlist) {

	conjContidos := []string{}

	for _, ranPlaylist := range conjPlaylists {

		for _, ranMusicas := range ranPlaylist.Musicas {

			if ranMusicas.Titulo == musica {

				conjContidos = append(conjContidos, ranPlaylist.Nome)

			}
		}
	}

	fmt.Printf("O nome das playlists que contêm a música '%s' são: ", musica)
	fmt.Print(conjContidos)

}

func playlistInfo(playlist Playlist) {

	soma := 0

	for _, ranMusicas := range playlist.Musicas {

		soma += ranMusicas.Duracao

	}

	fmt.Printf("Título: %s\n\nTempo total (em segundos): %d\n\nInformações das músicas:\n\n", playlist.Nome, soma)

	for _, ranMusicas := range playlist.Musicas {

		musicaInfo(ranMusicas)
		fmt.Printf("\n\n--------------------------------------------------------------------\n\n")

	}

}

func musicaInfo(musica Musica) {

	fmt.Println("Nome da música:", musica.Titulo)
	fmt.Println("Artista: ", musica.Artista)
	fmt.Printf("Duração da música (em segundos): %d\n\n", musica.Duracao)
	fmt.Print("Recomendação: ", musica.Recomendacao)
}

func printPlaylistsNames(conjPlaylists []Playlist) {

	for _, ranConjPlaylists := range conjPlaylists {

		fmt.Println(ranConjPlaylists.Nome)

	}

}

func menu() {

	fmt.Printf("\nDigite o número referente a o que você quer fazer:\n")
	fmt.Printf("\n1 - Informações de alguma playlist.\n2 - Editar alguma playlist.\n3 - Criar uma nova playlist. (não está funcionando)\n")
	fmt.Printf("4 - Apagar alguma playlist (não está funcionando)\n5 - Verificar em quais playlists está a sua música escolhida.\n")
	fmt.Printf("6 - Visualizar quais são as playlists. (não está funcionando)\nOutro - Encerrar o algoritmo.\n\n")
	fmt.Printf("----------------------------------------------------\n\n")

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	comando := ""

	PLMW := Musica{
		Titulo:       "Please Let Me Wonder",
		Artista:      "The Beach Boys",
		Duracao:      169,
		Recomendacao: "Gabriel",
	}

	B2 := Musica{
		Titulo:       "Bound 2",
		Artista:      "Kanye West",
		Duracao:      229,
		Recomendacao: "Alviner",
	}

	FILA := Musica{
		Titulo:       "Falling In Love Again",
		Artista:      "Joyce Manor",
		Duracao:      148,
		Recomendacao: "Gabriel",
	}

	SY := Musica{
		Titulo:       "Say Yes",
		Artista:      "Elliot Smith",
		Duracao:      140,
		Recomendacao: "Alviner",
	}

	japa := Musica{
		Titulo:       "Kimi no Shiranai",
		Artista:      "supercell",
		Duracao:      350,
		Recomendacao: "Breno",
	}

	BHolly := Musica{
		Titulo:       "Buddy Holly",
		Artista:      "Weezer",
		Duracao:      160,
		Recomendacao: "Mad Mordo",
	}

	CTD := Musica{
		Titulo:       "Cherish The Day",
		Artista:      "Sade",
		Duracao:      380,
		Recomendacao: "Mad e Alv",
	}

	conjMusicas1 := []Musica{PLMW, B2, FILA, SY, japa, BHolly, CTD}
	conjMusicas2 := []Musica{PLMW, B2}
	conjMusicas3 := []Musica{japa, BHolly, CTD}
	conjMusicas4 := []Musica{FILA, SY, japa}

	p1 := Playlist{
		Nome:    "7 músicas mais românticas.",
		Musicas: conjMusicas1,
	}

	p2 := Playlist{
		Nome:    "playlist 2",
		Musicas: conjMusicas2,
	}

	p3 := Playlist{
		Nome:    "playlist 3",
		Musicas: conjMusicas3,
	}

	p4 := Playlist{
		Nome:    "playlist 4",
		Musicas: conjMusicas4,
	}

	nomePlaylist := ""

	conjPlaylists := []Playlist{p1, p2, p3, p4}

	fmt.Println("Essas são as suas playlists:\n")
	printPlaylistsNames(conjPlaylists)
	fmt.Println("\n---------------------------------------")
	menu()

	scanner.Scan()
	comando = scanner.Text()
	fmt.Printf("\n")

	for {

		switch comando {

		case "1":

			fmt.Printf("Escreva o nome da playlist.\n\n")
			scanner.Scan()
			nomePlaylist = scanner.Text()
			fmt.Printf("\n")

			for _, ranConjPlaylists := range conjPlaylists {

				if nomePlaylist == ranConjPlaylists.Nome {

					playlistInfo(ranConjPlaylists)

				}

			}

			menu()

			scanner.Scan()
			comando = scanner.Text()
			fmt.Printf("\n")
			break

		case "2":

			fmt.Printf("Escreva o nome da playlist que você quer alterar.\n\n")
			scanner.Scan()
			selecionarPlaylist := scanner.Text()
			fmt.Printf("\n")

			fmt.Printf("Escolha o que você quer alterar:\n\n1- Nome da Playlist (não está funcionando)\n2- Música das Playlists\n\n")

			scanner.Scan()
			funcoes2 := scanner.Text()
			fmt.Printf("\n")

			if funcoes2 == "1" {

				fmt.Printf("Escreva o novo nome da sua playlist.\n\n")

				scanner.Scan()
				nomePlaylist = scanner.Text()
				fmt.Printf("\n")

				for _, ranConjPlaylists := range conjPlaylists {

					if selecionarPlaylist == ranConjPlaylists.Nome {

						ranConjPlaylists.Nome = nomePlaylist
						fmt.Printf("Nome alterado!\n\n")
						fmt.Printf("Agora essas são as suas playlists:\n\n")
						printPlaylistsNames(conjPlaylists)
						fmt.Println("\n---------------------------------------")

					}
				}

			}

			menu()

			scanner.Scan()
			comando = scanner.Text()
			fmt.Printf("\n")
			break

		case "5":

			fmt.Printf("Qual é música que você quer verificar?\n\n")

			scanner.Scan()
			nomeMusica := scanner.Text()
			fmt.Printf("\n")

			verifyTitles(nomeMusica, conjPlaylists)
			fmt.Printf("\n\n")

			menu()

			scanner.Scan()
			comando = scanner.Text()
			fmt.Printf("\n")
			break

		default:

			break

		}
		if comando != "1" && comando != "2" && comando != "3" && comando != "4" && comando != "5" && comando != "6" {
			break
		}
	}

}
