package main

import "fmt"

func ingresarTemp(temperaturas *[10]float64) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese la temperatura numero ", i+1)
		fmt.Scanln(&temperaturas[i])
	}
}

func main() {
	var temperaturas [10]float64
	var totalPacientes = float64(len(temperaturas))
	ingresarTemp(&temperaturas)

	var alta, normal, baja int

	for _, temp := range temperaturas {
		switch {
		case temp > 37.5:
			alta++
		case temp >= 36 && temp <= 37.5:
			normal++
		default:
			baja++
		}
	}

	porcentajeAlta := (float64(alta) / totalPacientes) * 100
	porcentajeNormal := (float64(normal) / totalPacientes) * 100
	porcentajeBaja := (float64(baja) / totalPacientes) * 100

	fmt.Println("Porcentaje con temperatura alta ", porcentajeAlta, "%")
	fmt.Println("Porcentaje con temperatura normal ", porcentajeNormal, "%")
	fmt.Println("Porcentaje con temperatura baja ", porcentajeBaja, "%")

	var (
		tmpMin = 999.0
		tmpMax = -1.0
	)

	for _, temp := range temperaturas {
		if temp > tmpMax {
			tmpMax = temp
		}
		if temp < tmpMin {
			tmpMin = temp
		}
	}
	fmt.Println("El promedio entre la temperatura maxima y la minima es: ", (tmpMax+tmpMin)/2)

}
