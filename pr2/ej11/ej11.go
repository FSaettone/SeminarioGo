package main

import (
	"fmt"
	"sort"
	"time"
)

type Ingresante struct {
	apellido        string
	nombre          string
	ciudadOrigen    string
	fechaNacimiento time.Time
	titulo          bool
	codigoCarrera   string
}

func (i Ingresante) String() string {
	str := fmt.Sprint("El ingresante " + i.apellido + " " + i.nombre + ", nacido en " + i.ciudadOrigen + " el " + i.fechaNacimiento.Format("02/01/2006"))
	if i.titulo {
		str += ", cuenta con el titulo secundario, y"
	} else {
		str += ", no cuenta con el titulo secundario, y"
	}
	str += fmt.Sprintln(" se encuentra inscripto en la carrera " + i.codigoCarrera)
	return str
}

func esMayor(a, b Ingresante) bool {
	return a.fechaNacimiento.Before(b.fechaNacimiento)
}

func compararNombre(i1, i2 Ingresante) bool {
	if i1.apellido == i2.apellido {
		return i1.nombre < i2.nombre
	}
	return i1.apellido < i2.apellido
}

func main() {
	fecha, _ := time.Parse("02/01/2006", "02/08/2001")
	fecha2, _ := time.Parse("02/01/2006", "14/06/2002")
	ingr := Ingresante{"Saettone", "Facundo", "Chascomus", fecha, true, "APU"}
	ingr2 := Ingresante{"Mariani", "Nicolas", "Chascomus", fecha2, false, "LS"}

	var ingresantes []Ingresante

	ingresantes = append(ingresantes, ingr)
	ingresantes = append(ingresantes, ingr2)

	fmt.Println("----------------ORDENADO POR EDAD----------------")
	sort.Slice(ingresantes, func(i, j int) bool {
		return esMayor(ingresantes[i], ingresantes[j])
	})

	for _, i := range ingresantes {
		fmt.Println(i)
	}
	fmt.Println("----------------ORDENADO ALFABETICAMENTE----------------")
	sort.Slice(ingresantes, func(i, j int) bool {
		return compararNombre(ingresantes[i], ingresantes[j])
	})

	for _, i := range ingresantes {
		fmt.Println(i)
	}
}
