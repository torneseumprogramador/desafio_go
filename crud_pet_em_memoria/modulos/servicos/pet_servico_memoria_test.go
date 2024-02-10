package servicos

import (
	"aula_go/crud_pet_em_memoria/modulos/enums"
	"aula_go/crud_pet_em_memoria/modulos/models"
	"testing"
)

func TestAdicionaPetCompleto(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = make([]models.Pet, 0)
	var pet models.Pet = models.Pet{
		Id:   1,
		Nome: "Lili",
		Dono: "Leandro",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	Adicionar(&pets, pet)

	// Assert (Verifica o resultado sobre o teste)
	if len(pets) == 0 {
		t.Errorf("Falha no test, deveria ter pelo menos 1 cadastrado na mamória")
	}
}

func TestAdicionaPetSemNome(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = make([]models.Pet, 0)
	var pet models.Pet = models.Pet{
		Id:   1,
		Nome: "",
		Dono: "Leandro",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Adicionar(&pets, pet)

	// Assert (Verifica o resultado sobre o teste)
	if erro == nil { // valido se não foi cadastrado
		t.Errorf("Não deveria deixar cadastrar um pet sem nome")
	} else { // valido a mensagem de erro
		mensagemEsperada := "O nome do pet é obrigatório"
		if mensagemEsperada != erro.Error() {
			t.Errorf("A mensagem deveria vir (%s) mas venho (%s)", mensagemEsperada, erro.Error())
		}
	}
}

func TestAdicionaPetBarraN(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = make([]models.Pet, 0)
	var pet models.Pet = models.Pet{
		Id:   1,
		Nome: "       \n",
		Dono: "Leandro",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Adicionar(&pets, pet)

	// Assert (Verifica o resultado sobre o teste)
	if erro == nil { // valido se não foi cadastrado
		t.Errorf("Não deveria deixar cadastrar um pet com \\n")
	} else { // valido a mensagem de erro
		mensagemEsperada := "O nome do pet é obrigatório"
		if mensagemEsperada != erro.Error() {
			t.Errorf("A mensagem deveria vir (%s) mas venho (%s)", mensagemEsperada, erro.Error())
		}
	}
}

func TestAdicionaPetSemDono(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = make([]models.Pet, 0)
	var pet models.Pet = models.Pet{
		Id:   1,
		Nome: "Lili",
		Dono: "",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Adicionar(&pets, pet)

	// Assert (Verifica o resultado sobre o teste)
	if erro == nil { // valido se não foi cadastrado
		t.Errorf("Não deveria deixar cadastrar um pet sem o dono")
	} else { // valido a mensagem de erro
		mensagemEsperada := "O dono do pet é obrigatório"
		if mensagemEsperada != erro.Error() {
			t.Errorf("A mensagem deveria vir (%s) mas venho (%s)", mensagemEsperada, erro.Error())
		}
	}
}

func TestAdicionaPetSemId(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = make([]models.Pet, 0)
	var pet models.Pet = models.Pet{
		Id:   0,
		Nome: "Lili",
		Dono: "Maria",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Adicionar(&pets, pet)

	// Assert (Verifica o resultado sobre o teste)
	if erro == nil { // valido se não foi cadastrado
		t.Errorf("Não deveria deixar cadastrar um pet sem ID")
	} else { // valido a mensagem de erro
		mensagemEsperada := "O ID de identificação, não pode ser <= 0"
		if mensagemEsperada != erro.Error() {
			t.Errorf("A mensagem deveria vir (%s) mas venho (%s)", mensagemEsperada, erro.Error())
		}
	}
}

func TestBuscarPorIdObjeto(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = listaCom2Pets()

	// Act (Roda a função que irá ser testada)
	pet, _ := BuscarPorId(pets, 2)

	// Assert (Verifica o resultado sobre o teste)
	if pet == nil {
		t.Errorf("Não foi encontrado o pet de ID 2 e deveria ser encontrado")
	}
}

func TestBuscarPorIdIndice(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = listaCom2Pets()

	// Act (Roda a função que irá ser testada)
	_, i := BuscarPorId(pets, 1)

	// Assert (Verifica o resultado sobre o teste)
	if i == -1 {
		t.Errorf("mandei localizar um pet válido porém retornou o indice de -1")
	}
}

func listaCom2Pets() []models.Pet {
	return []models.Pet{
		models.Pet{
			Id:   1,
			Nome: "Lili",
			Dono: "Maria",
			Tipo: enums.Gato,
		},
		models.Pet{
			Id:   2,
			Nome: "Josevaldo",
			Dono: "Geraldo",
			Tipo: enums.Cachorro,
		},
	}
}

func TestExcluir(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = listaCom2Pets()

	// Act (Roda a função que irá ser testada)
	pets, erro := Excluir(pets, 1)

	// Assert (Verifica o resultado sobre o teste)
	if erro != nil {
		t.Errorf("Erro ao excluir, não localizou o ID ou deu problema mesmo")
	}

	if len(pets) == 2 {
		t.Errorf("Erro, pois não excluiu da minha lista corretamente")
	}
}

func TestAlterar(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = listaCom2Pets()
	petAlteracao := models.Pet{
		Id:   1,
		Nome: "Lili a Gatinha",
		Dono: "Maria",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Alterar(&pets, petAlteracao)

	// Assert (Verifica o resultado sobre o teste)
	if erro != nil {
		t.Errorf("Erro ao fazer alteração do pet %s", erro.Error())
	}

	if pets[0].Nome != petAlteracao.Nome {
		t.Errorf("Erro, a alteração não foi aplicada da lista")
	}
}

func TestAlterarItemSemNome(t *testing.T) {
	// Arrange (prepara os itens para iniciar o teste)
	var pets []models.Pet = listaCom2Pets()
	petAlteracao := models.Pet{
		Id:   1,
		Nome: "",
		Dono: "Maria",
		Tipo: enums.Gato,
	}

	// Act (Roda a função que irá ser testada)
	erro := Alterar(&pets, petAlteracao)

	// Assert (Verifica o resultado sobre o teste)
	if erro == nil {
		t.Errorf("Erro, deixou alterar pet sem nome")
	}
}
