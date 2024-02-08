package main

import (
	"fmt"
	"strings"
)

func main() {

	// crio um map com chave e valor
	// pessoaMap := map[string]interface{}{
	// 	"Id":   2,
	// 	"Nome": "Bob",
	// }
	// fmt.Println(pessoaMap)

	// crio um map sem chave e valor
	pessoaMap := map[string]interface{}{}
	pessoaMap["Nome"] = "Danilo"         // adiciono o valor depois
	pessoaMap["Sobrenome"] = "Aparecido" // adiciono o valor depois
	fmt.Println(pessoaMap)

	// crio uma lista de map com itens
	var listaPessoasMap = []map[string]interface{}{
		{"Id": 1, "Nome": "Alice", "Cpf": "123413"},
		{"Id": 2, "Nome": "Bob"},
		{"Id": 3, "Nome": "Carlos"},
	}

	// var listaPessoasMap = []map[string]interface{}{} // Crio um map vazio

	// adiciono um item na lista de map
	listaPessoasMap = append(listaPessoasMap, map[string]interface{}{"Id": 4, "Nome": "Liah"})

	// // removo o ultimo item da lista de map
	// // Guarda o último item antes de fazer o pop
	// ultimoItem := listaPessoasMap[len(listaPessoasMap)-1] // pego o ultimo item da lista sem excluir

	// // Faz o pop do último item da lista de map
	// listaPessoasMap = listaPessoasMap[:len(listaPessoasMap)-1] // removo o ultimo item da lista e retorna a lista sem o ultimo item

	// fmt.Println("Nome:", ultimoItem["Nome"])

	// // como remover uma key do map
	// pessoaMap := listaPessoasMap[0]
	// fmt.Println("Antes de remover a chave CPF: ", pessoaMap)
	// delete(pessoaMap, "Cpf") // Remove a chave Cpf do map
	// fmt.Println("Depois de remover a chave CPF: ", pessoaMap)

	// // mostrando chave e valor do map
	// pessoaMap := listaPessoasMap[0]
	// for chave, valor := range pessoaMap {
	// 	fmt.Println("Chave:", chave, "Valor:", valor)
	// }

	fmt.Println(strings.Repeat("-", 40))

	for _, pessoa := range listaPessoasMap {
		fmt.Println("Id:", pessoa["Id"])
		fmt.Println("Nome:", pessoa["Nome"])
		valorDoCpf, existeChave := pessoa["Cpf"]
		if existeChave {
			fmt.Println("CPF:", valorDoCpf)
		} else {
			fmt.Println("CPF:", "Não cadastrado")
		}
		fmt.Println(strings.Repeat("-", 40))
	}

}
