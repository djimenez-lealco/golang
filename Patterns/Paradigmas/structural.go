package main

import "fmt"

type Cuadrado struct {
	lado int
}

func (cuadrado *Cuadrado) area() int {
	return cuadrado.lado * cuadrado.lado
}

func main() {
	fmt.Println("Structural")
	cuadrado := Cuadrado{lado: 4}
	// Al invocar el metodo de area se transfiere directamente el control al metodo especifico.
	fmt.Println(cuadrado.area())
	fmt.Println("Aqui se retorna el control")
}
