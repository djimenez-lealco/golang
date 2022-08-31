package main

import "fmt"

type IFigura interface {
	Area() int
}

type Cuadrado struct {
	lado int
}

func (cuadrado *Cuadrado) Area() int {
	return cuadrado.lado * cuadrado.lado
}

type Triangulo struct {
	base   int
	altura int
}

func (triangulo *Triangulo) Area() int {
	return (triangulo.base * triangulo.altura) / 2
}

func getFigura(figura string) IFigura {
	if figura == "triangulo" {
		return &Triangulo{base: 2, altura: 3}
	} else if figura == "cuadrado" {
		return &Cuadrado{lado: 4}
	} else {
		panic("Figura no reconocida")
	}
}

func imprimirFigura(figura IFigura) {
	// --------------------------------------------------------
	// Al invocar el metodo de area se transfiere
	// indirectamente el control al metodo Area de la
	// implementación segun la figura que corresponda.
	// --------------------------------------------------------
	fmt.Println(figura.Area())
}

func main() {
	fmt.Println("OOP")
	imprimirFigura(getFigura("triangulo"))
	imprimirFigura(getFigura("cuadrado"))
	fmt.Println("Aqui se retorna el control")
}
