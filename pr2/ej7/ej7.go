package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type crtrs struct {
	mayus      int
	minus      int
	nums       int
	especiales int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("ingrese una secuencia de caracteres")
	scanner.Scan()
	input := scanner.Text()
	var c crtrs
	contadorDigitos := make(map[rune]int)

	for _, char := range input {
		if unicode.IsLower(char) {
			c.minus++
		} else if unicode.IsUpper(char) {
			c.mayus++
		} else if unicode.IsDigit(char) {
			c.nums++
			contadorDigitos[char]++
		} else if unicode.IsSymbol(char) || unicode.IsSpace(char) {
			c.especiales++
		}
	}

	fmt.Println("la cantidad de letras minusculas detectadas es:", c.minus)
	fmt.Println("la cantidad de letras mayusculas detectadas es:", c.mayus)
	fmt.Println("la cantidad de numeros detectados es:", c.nums)
	fmt.Println("la cantidad de caracteres especiales detectados es:", c.especiales)

	fmt.Println("Ocurrencias de cada dígito:")
	for digito, count := range contadorDigitos {
		fmt.Printf("%c: %d\n", digito, count)
	}

}
