package condicionais

import "fmt"	

func IFs(){
	fmt.Println("IFs in Golang")

	// IF simples
	var i = 10
	fmt.Printf("Valor de i: %d\n", i)
	if i > 9{
		fmt.Println("i é maior que 9")
	}
	if i > 11{
		fmt.Println("i é maior que 11")
	}

	// IF com else
	if i > 11{
		fmt.Println("i é maior que 11")
	} else{
		fmt.Println("i é menor ou igual a 11")
	}

	// IF com else if
	if i > 11{
		fmt.Println("i é maior que 11")
	} else if i == 10{
		fmt.Println("i é igual a 10")
	} else{
		fmt.Println("i é menor que 10")
	}

	// IF com declaração
	if j := 10; j > 9{
		fmt.Println("j é maior que 9")
	}
 
	// if com declaração e else
	if j := 11; j%2 == 0{
		fmt.Println("j é par")
	} else{
		fmt.Println("j é impar")
	}


}

func SWITCHs(){
	fmt.Println("SWITCHs in Golang")
	var i int = 10
	switch{
	case i > 5:
		fmt.Println("i é maior que 5")
	case i > 10:
		fmt.Println("i é maior que 10")
	case i == 10:
		fmt.Println("i é igual a 10")
	default:
		fmt.Println("nenhuma das opções, mas o valor de i é: ", i)
	}

	// SWITCH com declaração
	name := "Mary"
	switch name {
	case "John":
		fmt.Println("Olá John")
	case "Mary":
		fmt.Println("Olá Mary")
	default:
		fmt.Println("Olá estranho")
	}
}