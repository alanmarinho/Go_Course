package structs

import (
	"fmt"
	"fundamentos/interfaces"
)

// struct é um tipo de dado que permite armazenar diferentes tipos de dados em um único tipo
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int16
	Salario   float64
}

// metodo de struct é uma função que pode ser chamada a partir de uma instância da struct
func (p Pessoa) ShowData() {
	fmt.Printf("Nome: %s %s\nIdade: %d\nSalário: %.2f\n", p.Nome, p.Sobrenome, p.Idade, p.Salario)
}

func (p Pessoa) ShowFullName() {
	fmt.Printf("Nome: %s %s\n", p.Nome, p.Sobrenome)
}

// esta garante que a struct Pessoa implementa a interface Pessoa, não é obrigatório, mas é uma boa prática
var _ interfaces.Pessoa = (*Pessoa)(nil)