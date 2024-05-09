package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ingresarTemp(temperaturas *[10]float64) {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese la temperatura numero ", i+1)
		scanner.Scan()
		temp, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("error al leer temperatura!!!")
			return
		}
		temperaturas[i] = temp
	}
}

func clasificarTemp(temperaturas []float64) {

	contadores := make(map[string]int)

	for i := range temperaturas {
		if temperaturas[i] >= 37.5 && temperaturas[i] < 50 {
			contadores["alta"] += 1
		} else if temperaturas[i] < 36 && temperaturas[i] > 20 {
			contadores["baja"] += 1
		} else if temperaturas[i] < 37.5 && temperaturas[i] >= 36 {
			contadores["normal"] += 1
		} else {
			contadores["incorrectos"] += 1
		}
	}
	total := 0
	for _, count := range contadores {
		total += count
	}

	porcentajeAlta := (float64(contadores["alta"]) / float64(total)) * 100
	porcentajeNormal := (float64(contadores["normal"]) / float64(total)) * 100
	porcentajeBaja := (float64(contadores["baja"]) / float64(total)) * 100

	fmt.Printf("Porcentaje de pacientes con temperatura alta: %.2f%%\n", porcentajeAlta)
	fmt.Printf("Porcentaje de pacientes con temperatura normal: %.2f%%\n", porcentajeNormal)
	fmt.Printf("Porcentaje de pacientes con temperatura baja: %.2f%%\n", porcentajeBaja)

	sort.Float64s(temperaturas[:])
	minTemp := temperaturas[0]
	maxTemp := temperaturas[len(temperaturas)-1]

	promedioInt := int((minTemp + maxTemp) / 2)

	fmt.Println("El promedio entero entre la temperatura minima:", minTemp, "y la temperatura maxima:", maxTemp, "es:", promedioInt)
}

func CelaFaren(temperaturas []float64) {

	for i, l := range temperaturas {
		temperaturas[i] = (l * 9 / 5) + 32
	}

	clasificarTemp(temperaturas[:])
}

func main() {
	var temperaturas [10]float64

	ingresarTemp(&temperaturas)
	clasificarTemp(temperaturas[:])
	fmt.Println("--------------------ahora en fahrenheit-----------------------------")
	CelaFaren(temperaturas[:])
}
