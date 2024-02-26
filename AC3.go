package main

import "fmt"

// Definindo o struct Contato
type Contato struct {
	Nome  string
	Email string
}

// Método para alterar o e-mail do Contato
func (c *Contato) alterarEmail(novoEmail string) {
	c.Email = novoEmail
}

// Função para adicionar um Contato ao array
func adicionarContato(contato Contato, arrayContatos *[5]Contato) {
	for i := range arrayContatos {
		if arrayContatos[i].Nome == "" {
			arrayContatos[i] = contato
			fmt.Println("Contato adicionado com sucesso!")
			return
		}
	}
	fmt.Println("O array de contatos está cheio. Não foi possível adicionar o contato.")
}

// Função para excluir o último Contato válido do array
func excluirContato(arrayContatos *[5]Contato) {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].Nome != "" {
			arrayContatos[i] = Contato{} // Limpa o contato
			fmt.Println("Último contato excluído com sucesso!")
			return
		}
	}
	fmt.Println("Nenhum contato válido para excluir.")
}

// Função para listar os contatos
func listarContatos(arrayContatos [5]Contato) {
	fmt.Println("\nLista de Contatos:")
	for i, contato := range arrayContatos {
		if contato.Nome != "" {
			fmt.Printf("%d. Nome: %s, Email: %s\n", i+1, contato.Nome, contato.Email)
		}
	}
}

func main() {
	// Criando um array de 5 elementos
	var contatos [5]Contato

	// Interface por linha de comando
	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Adicionar Contato")
		fmt.Println("2. Excluir Último Contato")
		fmt.Println("3. Listar Contatos")
		fmt.Println("4. Sair")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var novoContato Contato
			fmt.Println("Digite o nome do contato:")
			fmt.Scan(&novoContato.Nome)
			fmt.Println("Digite o e-mail do contato:")
			fmt.Scan(&novoContato.Email)
			adicionarContato(novoContato, &contatos)
		case 2:
			excluirContato(&contatos)
		case 3:
			listarContatos(contatos)
		case 4:
			fmt.Println("Saindo do programa.")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
