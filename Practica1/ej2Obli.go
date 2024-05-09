package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main2() {
	var frase string
	reader := bufio.NewReader((os.Stdin)) //creo un reader para leer un string con espacios incluidos

	fmt.Println("Ingrese una frase")

	frase, err := reader.ReadString('\n')

	frase = strings.TrimSpace(frase)
	if err == nil {
		if strings.Contains(strings.ToLower(frase), "miercoles") { //pasa la frase a minuscula y pregunta si contiene la palabra "miercoles"
			reemp2(frase)
		}
	} else {
		fmt.Println("no contiene miercoles")
	}

}

func reemp2(frase string) {
	fraseMod := ""
	indice := strings.Index(strings.ToLower(frase), "miercoles") //creo un indice a partir de la primera ocurrencia de "jueves"

	terminaMiercoles := 0 //creo un auxiliar

	for indice != -1 {
		fraseMod += frase[0:indice]   //corto y guardo la frase donde empieza la ocurrencia
		terminaMiercoles = indice + 9 //guardo la posicion de donde termina la palabra

		fraseMod += cambiarPalabra2(frase[indice:(indice + 9)]) //llamo al procedimiento que realiza el cambio respetando mayusculas o minusculas

		frase = frase[terminaMiercoles:] //reemplazo la frase original por la original cortada a partir de donde termina miercoles

		indice = strings.Index(strings.ToLower(frase), "miercoles") //actualizo el indice para buscar otra ocurrencia del miercoles
		if indice == -1 {
			fraseMod += frase[0:] //si no se encuentra la ocurrencia se copia lo que resta del string
		}
	}
	fmt.Println(fraseMod)
}

func cambiarPalabra2(s string) string {

	auxM := []rune(s)           // convierto el string de entrada en un slice de rune para acceder a cada caracter individualmente
	auxA := []rune("automovil") // lo mismo que el anterior pero para definir el string al que se le cambiaran los carateres

	for i, l := range auxM { // itera sobre cada caracter "l"
		if unicode.IsUpper(l) { // si el caracter es mayuscula, lo cambio por el almacenado en el slice convertido a mayuscula y viceversa
			auxA[i] = unicode.ToUpper(auxA[i])
		} else {
			auxA[i] = unicode.ToLower(auxA[i])
		}
	}
	return string(auxA)
}
