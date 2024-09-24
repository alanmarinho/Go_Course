package main

import (
	"fmt"
	"variaveis/exercicios"
)

func main(){
	temp1 := 37.5
	temp2 := 100.0
	temp3 := 0.0
	fmt.Println("Hello World!")
	fmt.Println(exercicios.CelciusToFahrenheit(temp1))
	fmt.Println(exercicios.CelciusToFahrenheit(temp2))
	fmt.Println(exercicios.CelciusToFahrenheit(temp3))

	fmt.Println(exercicios.FahrenheitToCelcius(temp1))
	fmt.Println(exercicios.FahrenheitToCelcius(temp2))
	fmt.Println(exercicios.FahrenheitToCelcius(temp3))
}