package main

import "fmt"

func ListarContatos(contatos []Contato) {
	if len(contatos) == 0 {
		fmt.Println("Nenhum contato cadastrado")
		return
	}

	fmt.Println("\nContatos cadastrados:")

	for i, contato := range contatos {
		fmt.Printf("\nContato %d\n", i+1)
		fmt.Println("Nome:", contato.Nome)
		fmt.Println("Telefone:", contato.Telefone)
		fmt.Println("Email:", contato.Email)
		fmt.Println("Observação:", contato.Observacao)
		fmt.Println("ID do contato:", contato.IdContato)
	}
}