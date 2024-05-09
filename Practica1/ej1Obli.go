package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main1() {
	var frase string
	reader := bufio.NewReader((os.Stdin)) //creo un reader para leer un string con espacios incluidos

	fmt.Println("Ingrese una frase")

	frase, err := reader.ReadString('\n')

	if err == nil {
		if strings.Contains(strings.ToLower(frase), "jueves") { //pasa la frase a minuscula y pregunta si contiene la palabra "jueves"
			reemp1(frase)
		}
	} else {
		fmt.Println("no contiene jueves")
	}

}

func reemp1(frase string) {
	fraseMod := " "
	indice := strings.Index(strings.ToLower(frase), "jueves") //creo un indice a partir de la primera ocurrencia de "jueves"

	terminaJueves := 0 //creo un auxiliar

	for indice != -1 {
		fraseMod += frase[0:indice] //corto y guardo la frase donde empieza la ocurrencia
		terminaJueves = indice + 6  //guardo la posicion de donde termina la palabra

		fraseMod += cambiarPalabra1(frase[indice:(indice + 6)]) //llamo al procedimiento que realiza el cambio respetando mayusculas o minusculas

		frase = frase[terminaJueves:] //reemplazo la frase original por la original cortada a partir de donde termina jueves

		indice = strings.Index(strings.ToLower(frase), "jueves") //actualizo el indice para buscar otra ocurrencia del jueves
		if indice == -1 {
			fraseMod += frase[0:] //si no se encuentra la ocurrencia se copia lo que resta del string
		}
	}
	fmt.Println(fraseMod)
}

func cambiarPalabra1(s string) string {

	auxJ := []rune(s)        // convierto el string de entrada en un slice de rune para acceder a cada caracter individualmente
	auxM := []rune("martes") // lo mismo que el anterior pero para definir el string al que se le cambiaran los carateres

	for i, l := range auxJ { // itera sobre cada caracter "l"
		if unicode.IsUpper(l) { // si el caracter es mayuscula, lo cambio por el almacenado en el slice convertido a mayuscula y viceversa
			auxM[i] = unicode.ToUpper(auxM[i])
		} else {
			auxM[i] = unicode.ToLower(auxM[i])
		}
	}
	return string(auxM)
}
