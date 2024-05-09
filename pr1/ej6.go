package main

import "fmt"

/*func main6() {
	var num1, num2 int

	fmt.Println("Ingresar un numero entero")
	fmt.Scanln(&num1)
	fmt.Println("Ingresar otro numero entero")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println("El resultado de dividir ", num1, " por ", num2, " es: ", num1/num2)
	}
	if num1 < num2 {
		fmt.Println("El resultado de dividir ", num2, " por ", num1, " es: ", num2/num1)
	}
}
*/
func main() {
	var num1, num2 float64

	fmt.Println("Ingresar un numero entero")
	fmt.Scanln(&num1)
	fmt.Println("Ingresar otro numero entero")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println("El resultado de dividir ", num1, " por ", num2, " es: ", num1/num2)
	}
	if num1 < num2 {
		fmt.Println("El resultado de dividir ", num2, " por ", num1, " es: ", num2/num1)
	}
}
