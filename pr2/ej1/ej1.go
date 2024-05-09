package main

import (
	"fmt"
	"sort"
)

func ingresarTemp(temperaturas *[10]float64) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese la temperatura numero ", i+1)
		fmt.Scanln(&temperaturas[i])
	}
}

func clasificarTemp(temperaturas []float64) {
	var alta, normal, baja int
	for i := range temperaturas {
		if temperaturas[i] > 37.5 {
			alta += 1
		} else if temperaturas[i] < 36 {
			baja += 1
		} else {
			normal += 1
		}
	}

	sort.Float64s(temperaturas[:])
	minTemp := temperaturas[0]
	maxTemp := temperaturas[len(temperaturas)-1]

	promedioInt := int((minTemp + maxTemp) / 2)

	fmt.Println("El promedio entero entre la temperatura minima ", minTemp, " y la temperatura maxima ", maxTemp, " es ", promedioInt)
}

func main() {
	var temperaturas [10]float64

	ingresarTemp(&temperaturas)
	clasificarTemp(temperaturas[:])
}
