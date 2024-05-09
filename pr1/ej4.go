package main

import "fmt"

const limite int = 0

func main() {
	for i := 250; i >= limite; i-- {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
