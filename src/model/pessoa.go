package model

//IPessoa - interface da pessoa
type IPessoa interface {
	Documento() string
}

//PessoaFisica - representa uma pessoa fisica
type PessoaFisica struct {
	Nome string
}

//PessoaJuridica - representa uma pessoa juridica
type PessoaJuridica struct {
	Nome string
}

//Documento - implementacao da pessoa para pessoa fisica
func (p PessoaFisica) Documento() string {
	return "CPF-->001.071.421-98"
}

//Documento - implementacao da pessoa para pessoa juridica
func (p PessoaJuridica) Documento() string {
	return "CNPJ-->26.885.644/0001-98"
}
