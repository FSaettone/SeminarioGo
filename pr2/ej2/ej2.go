package main

import "fmt"

func factorialIterativo(n int) int {
	if n < 0 {
		return -1 // Error: no se puede calcular el factorial de un número negativo
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
func factorialRecursivo(n int) int {
	if n < 0 {
		return -1 // Error: no se puede calcular el factorial de un número negativo
	}
	if n == 0 {
		return 1 // Caso base: 0! = 1
	}
	return n * factorialRecursivo(n-1) // Caso recursivo: n! = n * (n-1)!
}

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Printf("Factorial de %d (iterativo): %d\n", i, factorialIterativo(i))
	}
	for i := 0; i <= 9; i++ {
		fmt.Printf("Factorial de %d (recursivo): %d\n", i, factorialRecursivo(i))
	}
}
