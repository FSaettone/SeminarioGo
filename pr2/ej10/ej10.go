package main

import "fmt"

type Stack []int

func New() Stack {
	return make(Stack, 0)
}

func IsEmpty(s Stack) bool {
	return len(s) == 0
}

func Len(s Stack) int {
	return len(s)
}

func ToString(s Stack) string {
	str := "["
	for i, v := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(v)
	}
	str += "]"
	return str
}

func FrontElement(s Stack) int {
	if IsEmpty(s) {
		return 0
	}
	return s[len(s)-1]
}

func Push(s Stack, element int) Stack {
	return append(s, element)
}

func Pop(s Stack) (Stack, int) {
	if IsEmpty(s) {
		return s, 0
	}
	element := s[len(s)-1]
	return s[:len(s)-1], element
}

// Iterate aplica una función a cada elemento de la pila
func Iterate(s Stack, f func(int) int) {
	for _, v := range s {
		f(v)
	}
}

func main() {
	stack := New()
	fmt.Println("Pila vacía:", stack)

	stack = Push(stack, 1)
	stack = Push(stack, 2)
	stack = Push(stack, 3)
	fmt.Println("Pila después de pushear:", stack)

	stack, element := Pop(stack)
	fmt.Println("Elemento removido:", element)
	fmt.Println("Pila después de pop:", stack)

	fmt.Println("Longitud de la pila:", Len(stack))
	fmt.Println("Elemento al frente:", FrontElement(stack))

	Iterate(stack, func(v int) int {
		fmt.Println(v)
		return v
	})
}
