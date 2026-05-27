package main

import "fmt"

func ExcluirContato(contatos []Contato) []Contato {
	var nome string

	fmt.Print("Digite o nome do contato que deseja remover: ")
	fmt.Scanln(&nome)

	for i, contato := range contatos {
		if contato.Nome == nome {
			contatos = append(contatos[:i], contatos[i+1:]...)
			fmt.Println("Contato removido com sucesso")
			return contatos
		}
	}

	fmt.Println("Contato não encontrado")
	return contatos
}