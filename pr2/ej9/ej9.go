package main

type List = *nodo

type nodo struct {
	dato int
	sig  List
}

func New() List {
	var l List
	l = nil
	return l
}

func IsEmpty(self List) bool {
	if self == nil {
		return true
	} else {
		return false
	}
}

func Len(self List) int {
	count := 0
	for {
		if self != nil {
			count++
			self = self.sig
		} else {
			break
		}
	}
	return count
}

func FrontElement(self List) int {

	if self != nil {
		return self.dato
	} else {
		return -1
	}
}

func Next(self List) List {
	var l List
	l = nil
	if self != nil {
		return self.sig
	} else {
		return l
	}
}

func toString(self List) string {
	var a string
	return a
}

func main() {

}
