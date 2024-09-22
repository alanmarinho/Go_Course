package variables

import "fmt"

// escopo de pacote
var a = "Hello" // tipagem dinâmica
var b string = "World" // tipagem estática

// escopo de função
func escopo1() {
	fmt.Println("Escopo 1")
	var c int = 1
	fmt.Println(a, b, c)
}

func escopo2() {
	fmt.Println("Escopo 2")
	// fmt.Println(c) // c não está disponível nesse escopo
	var c int = 2  // c pode ser redeclarado nesse escopo
	fmt.Println(c)
}

func escopo3() {
	fmt.Println("Escopo 3")
	// variáveis podem ter valor atribuído em declaração
	var d int = 3
	fmt.Println(d)
	// variáveis podem ter valor atribuído após declaração
	d = 4
	fmt.Println(d)

	// variáveis podem ser declaradas sem valor
	var e int
	fmt.Println(e)

	// variáveis podem ser declaradas sem valor e atribuído depois
	var f int
	f = 5
	fmt.Println(f)

	// variáveis podem ser declaradas e atribuídas em uma única linha
	var g int = 6 // tipagem estática
	var h = 7 // tipagem inferida
	i := 8 // declaração curta

	fmt.Println(g, h, i)

	// somente o simbolo de atribuição = pode ser usado para atribuir valor a uma variável
	g = 9 // ok
	// := é usado para declarar e atribuir valor a uma variável pela primeira vez
	// g := 9 // erro

	fmt.Println(g)

}


// se a função for exportada, ela deve começar com letra maiúscula (Escopo)
// se a função não for exportada, ela deve começar com letra minúscula (escopo1 e escopo2) e só será visível dentro do pacote
func Escopo() {
	escopo1()
	escopo2()
	escopo3()
}

