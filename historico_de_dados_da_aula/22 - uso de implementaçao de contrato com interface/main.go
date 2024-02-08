package main

import (
	"fmt"
)

type InterfacePessoa interface {
	ValidarDocumento() bool
	Documento() string
}

type PessoaFisica struct {
	Nome string
	CPF  string
}

func (pessoaFisica PessoaFisica) ValidarDocumento() bool {
	if pessoaFisica.CPF != "" {
		return true
	}
	return false
}

func (pessoaFisica PessoaFisica) Documento() string {
	return pessoaFisica.CPF
}

type PessoaJurifica struct {
	Nome string
	CNPJ string
}

func (pessoaJuridica PessoaJurifica) ValidarDocumento() bool {
	if pessoaJuridica.CNPJ != "" {
		return true
	}
	return false
}

func (pessoaJuridica PessoaJurifica) Documento() string {
	return pessoaJuridica.CNPJ
}

func main() {
	pessoaFisica := new(PessoaFisica)
	pessoaFisica.Nome = "Danilo"
	pessoaFisica.CPF = "377.483.620-50"

	pessoaJurifica := new(PessoaJurifica)
	pessoaJurifica.Nome = "D&D"
	pessoaJurifica.CNPJ = "97.793.972/0001-12"

	VarificaSeDocumentoEstaOk(pessoaFisica)
	VarificaSeDocumentoEstaOk(pessoaJurifica)
}

func VarificaSeDocumentoEstaOk(iPessoa InterfacePessoa) {
	if iPessoa.ValidarDocumento() {
		fmt.Printf("O documento (%T) %s não está vazio\n", iPessoa, iPessoa.Documento())
	} else {
		fmt.Printf("O documento (%T) %s está vazio\n", iPessoa, iPessoa.Documento())
	}
}
