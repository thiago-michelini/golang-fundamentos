package model

//Cliente - struct que representa cliente
type Cliente struct {
	Nome     string `json:"nomeCli"`
	Cpf      string `json:"cpfCli"`
	Idade    int
	salario  float64
	Endereco *Endereco
}

//SetSalario - set salario
func (cliente *Cliente) SetSalario(sal float64) {
	cliente.salario = sal
}

//GetSalario - get salario
func (cliente *Cliente) GetSalario() float64 {
	return cliente.salario
}
