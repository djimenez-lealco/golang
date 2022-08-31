package main

import (
	"fmt"

	saludar "identity.leal.co/infraestructure/Saludar"
)

func main() {
	fmt.Println("hola mundo")
	fmt.Println(saludar.SaludarFactory().Saludar())
}
