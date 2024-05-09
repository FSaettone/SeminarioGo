package main

import "fmt"

func main() {
	var ptoCardinal string

	fmt.Println("Ingrese un punto cardinal N,S,E,O")
	fmt.Scanln(&ptoCardinal)

	switch ptoCardinal {
	case "N":
		fmt.Println("El viento se dirije al Norte")
	case "S":
		fmt.Println("El viento se dirije al Sur")
	case "E":
		fmt.Println("El viento se dirije al Este")
	case "O":
		fmt.Println("El viento se dirije al Oeste")
	default:
		fmt.Println("Entrada no valida")
	}
}
