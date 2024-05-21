package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
	"time"
)

type ingresante struct {
	apellido        string
	nombre          string
	ciudadOrigen    string
	fechaNacimiento time.Time
	titulo          bool
	codigoCarrera   string
}

func procesarIngresantes(lista *list.List) {
	bariloche := []string{}
	contadorNacimientos := make(map[int]int)

	añoMax := 0
	cantAñoMax := 0

	contadorCarreras := make(map[string]int)
	carreraMax := ""
	cantCarreraMax := 0

	i := lista.Front()

	for i != nil {
		ingresante := i.Value.(ingresante)
		siguiente := i.Next()

		if ingresante.ciudadOrigen == "Bariloche" {
			bariloche = append(bariloche, fmt.Sprint(ingresante.apellido, " ", ingresante.nombre))
		}

		año := ingresante.fechaNacimiento.Year()
		fmt.Println(" año actual", año)
		contadorNacimientos[año]++
		if contadorNacimientos[año] > cantAñoMax {
			añoMax = año
			cantAñoMax = contadorNacimientos[año]
		}

		contadorCarreras[ingresante.codigoCarrera]++
		if contadorCarreras[ingresante.codigoCarrera] > cantCarreraMax {
			cantCarreraMax = contadorCarreras[ingresante.codigoCarrera]
			carreraMax = ingresante.codigoCarrera
		}

		if !ingresante.titulo {
			lista.Remove(i)
			fmt.Println("Se elimino el ingresante", ingresante.apellido, "por no tener el titulo secundario")
		}
		i = siguiente
	}

	fmt.Println("Ingresantes provenientes de Bariloche:")
	for _, nombre := range bariloche {
		fmt.Println(nombre)
	}

	fmt.Println("El año con mas nacimientos es el ", añoMax, " con ", cantAñoMax, "ingresantes")
	fmt.Println("La carrera con mas inscriptos es ", carreraMax)

}

func main() {
	texto, err := os.Open("ingresantes.txt")
	if err != nil {
		return
	}
	defer texto.Close()

	lista := list.New()

	scanner := bufio.NewScanner(texto)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) == 6 {
			nombre := fields[0]
			apellido := fields[1]
			ciudadOrigen := fields[2]
			fechaNacimiento, _ := time.Parse("01/02/2006", fields[3])
			titulo := fields[4] == "true"
			codigoCarrera := fields[5]

			ingresante := ingresante{
				apellido:        apellido,
				nombre:          nombre,
				ciudadOrigen:    ciudadOrigen,
				fechaNacimiento: fechaNacimiento,
				titulo:          titulo,
				codigoCarrera:   codigoCarrera,
			}
			lista.PushBack(ingresante)
			fmt.Println("se cargo un ingresante", ingresante.nombre, " ", ingresante.fechaNacimiento.Year())
		}
	}
	procesarIngresantes(lista)

	//Get-Content ingresantes.txt | go run ej1.go
}
