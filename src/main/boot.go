package main

import (
	"encoding/json"
	"fmt"

	"../model"
)

func main() {
	//**************************************************************************************
	//*********************Struct (equivalente a Classes no java)***************************
	//**************************************************************************************
	cliente := new(model.Cliente)
	cliente.Nome = "thiago"
	cliente.Cpf = "00107142198"
	cliente.Endereco = new(model.Endereco)
	cliente.Endereco.Cep = "79115530"
	cliente.Endereco.Latitude = "515496546-58"
	cliente.Endereco.Longitude = "515496546-58"
	fmt.Printf("Cliente é --> %+v \n", cliente)

	cliente2 := new(model.Cliente)
	cliente2.Nome = "teste"
	cliente2.Cpf = "9999999"
	cliente2.SetSalario(99.6)
	cliente2.Idade = 25
	fmt.Printf("Cliente2 é --> %+v \n", cliente2)

	//**************************************************************************************
	//*****************************************Maps*****************************************
	//**************************************************************************************
	mapa := make(map[string]*model.Cliente)
	mapa[cliente.Nome] = cliente
	mapa[cliente2.Nome] = cliente2
	fmt.Printf("mapa é --> %+v \n", mapa)

	//**************************************************************************************
	//**********************Slices (equivalente a List no java)*****************************
	//**************************************************************************************
	lista := make([]*model.Cliente, 0)
	lista = append(lista, cliente)
	lista = append(lista, cliente2)
	for idx, cli := range lista {
		fmt.Printf("indice[%d] --> valor %+v\n", idx, cli)
	}

	//**************************************************************************************
	//***********************************Parses com JSON************************************
	//**************************************************************************************
	strJSON, _ := json.Marshal(cliente)
	fmt.Printf("struct to json --> %s\n", string(strJSON))
	cliente.Nome = cliente.Nome + " (json foi alterado)"
	cliente.Cpf = cliente.Cpf + " (json foi alterado)"
	strJSON, _ = json.Marshal(cliente)

	var cliJSON model.Cliente
	_ = json.Unmarshal([]byte(strJSON), &cliJSON)
	fmt.Printf("json to struct --> nome: %s | cpf: %s\n", cliJSON.Nome, cliJSON.Cpf)

	//**************************************************************************************
	//*********Polimorfismo (no Go só existe interfaces, não há herança explicita)**********
	//**************************************************************************************
	var tpPessoa string
	fmt.Println("*************************************************\nInforme o tipo de pessoa (F ou J):")
	fmt.Scanln(&tpPessoa)
	var pessoa model.IPessoa
	if tpPessoa == "J" {
		pessoa = model.PessoaJuridica{Nome: "Nome da pessoa juridica"}
	} else {
		pessoa = model.PessoaFisica{Nome: "Nome da pessoa fisica"}
	}
	fmt.Printf("Documento.: %s, com a implementação: %T\n", pessoa.Documento(), pessoa)
}
