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
