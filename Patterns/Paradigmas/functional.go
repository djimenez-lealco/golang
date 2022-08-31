package main

import "fmt"

func Close(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
func main() {
	// CLOSURES

	fmt.Println("Programacion Funcional en Go")
	sumador := Close(2)
	fmt.Println(sumador(3))
}
