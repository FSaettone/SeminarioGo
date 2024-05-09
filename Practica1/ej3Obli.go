package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main3() {

	reader := bufio.NewReader((os.Stdin)) //creo un reader para leer un string con espacios incluidos

	fmt.Println("Ingrese una frase")

	frase, err := reader.ReadString('\n')
	fmt.Println("Ingrese una palabra")
	var palabra string
	fmt.Scanln(&palabra)

	if err == nil {
		if strings.Contains(strings.ToLower(frase), palabra) { //pasa la frase a minuscula y pregunta si contiene la palabra
			reemp3(frase, palabra)
		}
	} else {
		fmt.Println("no contiene " + palabra)
	}

}

func reemp3(frase string, palabra string) {
	fraseMod := ""
	indice := strings.Index(strings.ToLower(frase), palabra) //creo un indice a partir de la primera ocurrencia de la palabra

	termina := 0

	for indice != -1 {
		fraseMod += frase[0:indice]     //corto y guardo la frase donde empieza la ocurrencia
		termina = indice + len(palabra) //guardo la posicion de donde termina la palabra

		fraseMod += cambiarPalabra3(frase[indice:(indice + len(palabra))]) //llamo al procedimiento que realiza el cambio respetando mayusculas o minusculas

		frase = frase[termina:] //reemplazo la frase original por la original cortada a partir de donde termina la palabra

		indice = strings.Index(strings.ToLower(frase), palabra) //actualizo el indice para buscar otra ocurrencia
		if indice == -1 {
			fraseMod += frase[0:] //si no se encuentra la ocurrencia se copia lo que resta del string
		}
	}
	fmt.Println(fraseMod)
}

func cambiarPalabra3(s string) string {

	aux := []rune(s) // convierto el string de entrada en un slice de rune para acceder a cada caracter individualmente

	for i, l := range aux { // itera sobre cada caracter "l"
		if unicode.IsUpper(l) { // si el caracter es mayuscula, lo invierto a minuscula y viceversa
			aux[i] = unicode.ToLower(aux[i])
		} else {
			aux[i] = unicode.ToUpper(aux[i])
		}
	}
	return string(aux)
}
