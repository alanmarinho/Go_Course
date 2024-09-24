package exercicios

import (
	"fmt"
)

func CelciusToFahrenheit(temerature float64) string {
	var fahrenheit float64 = temerature * 1.8 + 32
	return fmt.Sprintf("%.2f°F", fahrenheit) // Detalhe, o SprintF serve para retornar uma string formatada não para imprimir na tela
}

func FahrenheitToCelcius(temperature float64) string{
	var celcius float64 = (temperature - 32) / 1.8
	return fmt.Sprintf("%.2f°C", celcius) 
}
