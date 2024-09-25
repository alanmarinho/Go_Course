package labels

import "fmt"

// Labels in Golang - Labels são usados para controlar o fluxo de execução de um programa
// é desecorado o uso de labels de Go pois almenta a complexidade do fluxo do código
// oq é o oposto do proposto pela linguagem que é ser simples e direta

func Labels(){
	var i = 0
	label1:
		fmt.Println("Label 1")
		i++
		goto label4
		
	label2:
		fmt.Println("Label 2")
		i++
		if i > 8{
			goto label1
		} 
		goto label3

	label3:
		fmt.Println("Label 3")
		i++
		goto label2

	label4:
		fmt.Println("Label 4")
		if i < 10{
			goto label2
		} else {
			exit()
		}
}

func exit(){
	fmt.Println("Exit")
}