package main

import (
	"desafios/escalas"
	"fmt"
)

func main(){
	var ebulicaoAguaKelvin float64 = 373.15
	fmt.Println("Ponto de ebulição da água de Kelvin para Celsius")
	fmt.Printf("Ponto de ebulição da água: %.2f°C", escalas.KelvinToCelsius(ebulicaoAguaKelvin))
}