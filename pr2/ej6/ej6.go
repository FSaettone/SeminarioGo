package main

import "fmt"

func Sum(a, b []int) []int {
	tamaño := len(a)
	if len(b) < tamaño {
		tamaño = len(b)
	}
	res := make([]int, tamaño)

	for i := range tamaño {
		res[i] = a[i] + b[i]
	}
	return res
}

func AvgInt(a []int) int {
	var res int
	for _, l := range a {
		res += l
	}
	return res / len(a)
}
func AvgFloat(a []int) float64 {
	var res int
	for _, l := range a {
		res += l
	}
	return float64(res) / float64(len(a))
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}

	fmt.Println("La suma de los slices da como resultado  ", Sum(s1, s2))

	fmt.Println("el promedio de el slice 1 da ", AvgInt(s1))
	fmt.Println("el promedio de el slice 2 da ", AvgInt(s2))

	fmt.Println("el promedio de el slice 1 da ", AvgFloat(s1))
	fmt.Println("el promedio de el slice 2 da ", AvgFloat(s2))
}
