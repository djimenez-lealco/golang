package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre string
	Color  string
	Edad   int
}

func main() {
	p1 := Persona{Nombre: "Daniel", Color: "Verde", Edad: 36}

	tipo := reflect.TypeOf(p1)
	nombre := tipo.Name()
	clase := tipo.Kind()
	valor := reflect.ValueOf(p1)

	pp1 := &p1
	fmt.Println(nombre, tipo, clase, valor)

	gType := reflect.TypeOf(p1)
	fmt.Printf("Type: %s Kind: %s\n", gType.Name(), gType.Kind())

	strNumFields := gType.NumField()
	for i := 0; i < strNumFields; i++ {
		field := gType.Field(i)
		fmt.Printf("Field Type: %s: %s Kind: %s\n", field.Name, field.Type.Name(), field.Type.Kind())
	}

	ptrType := reflect.TypeOf(pp1)
	fmt.Printf("Type: %s Kind: %s\n", ptrType.Elem(), ptrType.Kind())

}
