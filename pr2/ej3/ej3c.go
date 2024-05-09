package main

import "fmt"

func main() {
	puntosCardinales := map[string]string{
		"N":  "S",
		"S":  "N",
		"E":  "O",
		"O":  "E",
		"NO": "SE",
		"NE": "SO",
		"SE": "NO",
		"SO": "NE",
	}

	var input string

	fmt.Println("Ingrese la direccion del viento (N,S,E,O,NO,NE,SO,SE)")
	fmt.Scanln(&input)

	direcOpuesta, ok := puntosCardinales[input]
	if !ok {
		fmt.Println("Direccion invalida")
		return
	}

	fmt.Println("si el viento viene del", input, ", se dirije al", direcOpuesta)
}
