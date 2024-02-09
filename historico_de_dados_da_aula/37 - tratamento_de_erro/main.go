package main

import (
	"errors"
	"fmt"
	"reflect"
)

// CampoVazioError representa um erro para campos vazios.
type CampoVazioError struct {
	Campo string
}

func (e *CampoVazioError) Error() string {
	return fmt.Sprintf("O campo '%s' está vazio", e.Campo)
}

// CampoUnicoErros representa um erro para campos que deveriam ser únicos mas não são.
type CampoUnicoErros struct {
	Campo string
}

func (e *CampoUnicoErros) Error() string {
	return fmt.Sprintf("O campo '%s' deve ser único", e.Campo)
}

// AlgumaOperacao simula uma operação que pode gerar diferentes tipos de erros.
func AlgumaOperacao(i int) error {
	switch i {
	case 1:
		return &CampoVazioError{"Nome"}
	case 2:
		return &CampoUnicoErros{"Email"}
	default:
		return errors.New("Erro genérico")
	}
}

func mostra_mensagem(erro error) {
	tipoErro := reflect.TypeOf(erro)
	if tipoErro.Kind() == reflect.Ptr {
		tipoErro = tipoErro.Elem()
	}
	fmt.Printf("Erro de %s: %s\n", tipoErro, erro.Error())
}

func main() {
	for i := 1; i <= 3; i++ {
		err := AlgumaOperacao(i)
		mostra_mensagem(err)

		//======== condicional switch ========
		// switch e := err.(type) {
		// case *CampoVazioError:
		// 	fmt.Println("Erro de Campo Vazio:", e.Error())
		// case *CampoUnicoErros:
		// 	fmt.Println("Erro de Campo Único:", e.Error())
		// default:
		// 	fmt.Println("Erro genérico:", err.Error())
		// }

		//======== condicional if ========
		// if _, ok := err.(*CampoVazioError); ok {
		// 	fmt.Printf("Erro validado CampoVazioError: %s\n", err.Error())
		// } else if _, ok := err.(*CampoUnicoErros); ok {
		// 	fmt.Printf("Erro validado CampoUnicoErros: %s\n", err.Error())
		// } else
		// 	fmt.Println("Erro genérico:", err.Error())
		// }
	}
}
