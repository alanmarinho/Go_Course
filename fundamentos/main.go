package main

import (
	"fundamentos/structs"
	"fundamentos/teste"
	"fundamentos/variables"
)



func main(){

	teste.Hello()
	variables.Escopo()
	pessoa := structs.Pessoa{Nome: "João", Sobrenome: "Silva", Idade: 30, Salario: 1000.0}
	pessoa.ShowData()
	pessoa.ShowFullName()

}