package main

import "fmt"

type Pessoa struct {
	Id   int
	Nome string
}

var listaPessoasStruct = []Pessoa{
	{Id: 1, Nome: "Alice"},
	{Id: 2, Nome: "Bob"},
	{Id: 3, Nome: "Carlos"},
}
var listaPessoasMap = []map[string]interface{}{
	{"Id": 1, "Nome": "Alice", "Cpf": "123413"},
	{"Id": 2, "Nome": "Bob"},
	{"Id": 3, "Nome": "Carlos"},
}
var listaPessoasArrays = [][2]string{
	{"1", "Alice"},
	{"2", "Bob"},
	{"3", "Carlos"},
}

func main() {
	listaPessoasArrays = append(listaPessoasArrays, [2]string{"4", "Liah"})
	listaPessoasStruct = append(listaPessoasStruct, Pessoa{Id: 4, Nome: "Liah"})
	listaPessoasMap = append(listaPessoasMap, map[string]interface{}{"Id": 4, "Nome": "Liah"})

	nomeDoArray := listaPessoasArrays[0][1]
	fmt.Println("Nome:", nomeDoArray)

	nomeDoMap := listaPessoasMap[0]["Nome"]
	fmt.Println("Nome:", nomeDoMap)

	nomeDoStruct := listaPessoasStruct[0].Nome
	fmt.Println("Nome:", nomeDoStruct)

	for i, pessoa := range listaPessoasArrays {
		fmt.Println("Indice:", i)
		fmt.Println("Id:", pessoa[0])
		fmt.Println("Nome:", pessoa[1])
	}

	for i, pessoa := range listaPessoasMap {
		fmt.Println("Indice:", i)
		fmt.Println("Id:", pessoa["Id"])
		fmt.Println("Nome:", pessoa["Nome"])
	}

	for i, pessoa := range listaPessoasStruct {
		fmt.Println("Indice:", i)
		fmt.Println("Id:", pessoa.Id)
		fmt.Println("Nome:", pessoa.Nome)
	}

	// var meuMap map[string]int
	// meuMap = make(map[string]int)

	// meuMap["chave1"] = 10
	// meuMap["chave2"] = 20

	// fmt.Printf("%d", meuMap["chave2"])

	// var cofre = Cofre{
	// 	Chave1: 10,
	// 	Chave2: 20,
	// }

	// fmt.Printf("%d", cofre.Chave2)
}
