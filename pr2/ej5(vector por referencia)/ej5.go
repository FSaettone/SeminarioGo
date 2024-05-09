package main

import "fmt"

const dimF = 3

type Vector [dimF]float64

func initialize(v *Vector, f float64) {

	for i := range v {
		v[i] = f
	}

}

func Sum(v1, v2 Vector) Vector {
	var suma Vector

	for i := range v1 {
		suma[i] = v1[i] + v2[i]
	}
	return suma
}

func SumInPlace(v1, v2 *Vector) {
	for i := range v1 {
		v1[i] += v2[i]
	}
}

func main() {
	v1 := Vector{0.0, 1.1, 2.2}
	v2 := Vector{3.3, 4.4, 5.5}
	var v3 Vector

	initialize(&v3, 5.5)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	fmt.Println("Sum de v1 y v2 ", Sum(v1, v2))
	SumInPlace(&v3, &v2)
	fmt.Println("SumInPlace de v3 y v2", v3)

}
