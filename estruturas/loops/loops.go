package loops

import "fmt"

// em Go só existe o for como Loop. Não existe while, do while, etc.
func ForLoops(){
	fmt.Println("For loops in Golang")
	
	// For loop
	for i := 0; i < 2; i++ {
		fmt.Printf("Contador: %d\n", i)
	}

	// for em listas
	var list = []string{"a", "b", "c"}
	for i, v := range list{
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	// for em mapas
	var m = map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m{
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	// for condicional (substitui while)
	var i = 3
	for i > 0 {
		fmt.Printf("Loop condicional %d\n", i)
		i--
	}

	// for limpo com break (alto risco de loop infinito - cuidado) - O brake é usado para sair do loop completamente
	for {
		fmt.Printf("Loop infinito com break %d\n", i)
		i++
		if i > 5 {
			break
		}

	}

	// for com continue  (o continue pula a iteração atual e vai para a próxima iteração sem quebrar o loop)
	for i := 0; i < 5; i++ {
		if i == 3{
			continue
		}
		fmt.Printf("Contador: %d\n", i)
	}

	


	
	// For loop sem condição
	// for {
	// 	fmt.Println("Loop infinito")	
	// }

	// for var i int = 0; i < 2: i++{  // somente declaração com operador curto :=
	// 	fmt.Printf("Contador: %d\n", i)
	// }




}