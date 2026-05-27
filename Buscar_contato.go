package main

import "fmt"

func BuscarContato(contatos []Contato) []Contato {
	var nome string

	fmt.Print("Digite o nome do contato que deseja buscar: ")
	fmt.Scanln(&nome)

	var encontrados []Contato

	for i, contato := range contatos {
		if contato.Nome == nome {
			fmt.Printf("Contato encontrado na posição %d\n", i)
			encontrados = append(encontrados, contato)
		}
	}

	if len(encontrados) == 0 {
		fmt.Println("Contato não encontrado")
	}

	return encontrados
}