package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CadastrarContato() Contato {
	var contato Contato

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o ID do contato: ")
	fmt.Scanln(&contato.IdContato)
	reader.ReadString('\n')

	fmt.Print("Digite o nome do contato: ")
	contato.Nome, _ = reader.ReadString('\n')
	contato.Nome = strings.TrimSpace(contato.Nome)

	fmt.Print("Digite o telefone: ")
	contato.Telefone, _ = reader.ReadString('\n')
	contato.Telefone = strings.TrimSpace(contato.Telefone)

	fmt.Print("Digite o email: ")
	contato.Email, _ = reader.ReadString('\n')
	contato.Email = strings.TrimSpace(contato.Email)

	fmt.Print("Digite uma observação: ")
	contato.Observacao, _ = reader.ReadString('\n')
	contato.Observacao = strings.TrimSpace(contato.Observacao)

	return contato
}