package main

import "fmt"

func main() {
	contatos := []Contato{}

	for {
		fmt.Println("\n===== Bem-vindo a agenda de contatos =====")
		fmt.Println("1 - Cadastrar contato")
		fmt.Println("2 - Listar contatos")
		fmt.Println("3 - Buscar contato")
		fmt.Println("4 - Remover contato")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int

		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Opção inválida. Digite apenas números.")
			fmt.Scanln()
			continue
		}

		switch opcao {
		case 1:
			novoContato := CadastrarContato()
			contatos = append(contatos, novoContato)
			fmt.Println("Contato cadastrado com sucesso")

		case 2:
			ListarContatos(contatos)

		case 3:
			encontrados := BuscarContato(contatos)
			ListarContatos(encontrados)

		case 4:
			contatos = ExcluirContato(contatos)

		case 0:
			fmt.Println("Encerrando sistema...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}