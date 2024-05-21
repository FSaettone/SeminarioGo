package main

import "fmt"

type contador map[int]int

type OptimumSlice []contador

func New(slice []int) OptimumSlice {
	var o OptimumSlice
	var contadorAct contador
	var ant int

	for _, act := range slice {
		if contadorAct == nil {
			contadorAct = make(contador)
			ant = act
		}
		if act == ant {
			contadorAct[act]++
		} else {
			o = append(o, contadorAct)
			contadorAct = make(contador)
			contadorAct[act]++
			ant = act
		}
	}
	if contadorAct != nil {
		o = append(o, contadorAct)
	}
	return o

}

func (o *OptimumSlice) isEmpty() bool {
	if len(*o) == 0 {
		return true
	}
	return false
}

func (o *OptimumSlice) Len() int {
	return len(*o)
}

func (o *OptimumSlice) FrontElement() int {
	if o.isEmpty() {
		return 0
	}
	valor := *o
	for e := range valor[0] {
		return e
	}
	return 0
}

func (o *OptimumSlice) LastElement() int {
	if o.isEmpty() {
		return 0
	}
	valor := *o
	for e := range valor[len(valor)-1] {
		return e
	}
	return 0
}

func (o *OptimumSlice) Insert(element, position int) int {
	if position < 0 || position > o.Len() {
		return -1
	}

	return 0
}

func (o *OptimumSlice) SliceArray() []int {
	var s []int

	for _, contador := range *o {
		for clave, valor := range contador {
			for i := 0; i < valor; i++ {
				s = append(s, clave)
			}
		}
	}
	return s
}

func main() {
	var o OptimumSlice
	fmt.Println(o.isEmpty())
	slice := []int{1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 3, 3, 2, 2, 2}
	o = New(slice)
	fmt.Println(o)
	fmt.Println(o.isEmpty())
	fmt.Println(o.Len())
	fmt.Println(o.FrontElement())
	fmt.Println(o.LastElement())
	fmt.Println(o.SliceArray())

}
