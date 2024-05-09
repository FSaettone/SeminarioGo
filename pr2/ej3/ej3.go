package main

import "fmt"

type PuntoCardinal int

const (
	Norte PuntoCardinal = iota
	Sur
	Este
	Oeste
	NorEste
	NorOeste
	SurEste
	SurOeste
)

func obtainDirecOpuesta(direcViento PuntoCardinal) string {
	switch direcViento {
	case Norte:
		return "Sur"
	case Sur:
		return "Norte"
	case Este:
		return "Oeste"
	case Oeste:
		return "Este"
	case NorEste:
		return "SurEste"
	case NorOeste:
		return "SurOeste"
	case SurEste:
		return "NorOeste"
	case SurOeste:
		return "NorEste"
	default:
		return "Direccion invalida"
	}
}

func main1() {
	var direcViento PuntoCardinal
	var input string

	fmt.Println("Ingrese la direccion del viento (N,S,E,O,NO,NE,SO,SE)")
	fmt.Scanln(&input)

	switch input {
	case "N":
		direcViento = Norte
	case "S":
		direcViento = Sur
	case "E":
		direcViento = Este
	case "O":
		direcViento = Oeste
	case "NO":
		direcViento = NorOeste
	case "NE":
		direcViento = NorEste
	case "SO":
		direcViento = SurOeste
	case "SE":
		direcViento = SurEste
	default:
		fmt.Println("Direccion no valida!")
	}

	direcOpuesta := obtainDirecOpuesta(direcViento)

	fmt.Println("si el viento viene del", input, ", se dirije al", direcOpuesta)
}
