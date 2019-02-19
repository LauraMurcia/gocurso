package main

import (
	"fmt"

	"github.com/LauraMurcia/gocurso/adivinanumero"
	"github.com/LauraMurcia/gocurso/calculadora"
	"github.com/LauraMurcia/gocurso/holamundo"
	"github.com/LauraMurcia/gocurso/maps"
	"github.com/LauraMurcia/gocurso/structs"
)

func main() {
	calculadora.Calcular()
	adivinanumero.Adivina()

	nombre := holamundo.GetName()
	fmt.Println("Nombre ingresado: ", nombre)

	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetValorMap("Laura"))

	structs.InterfaceTest()
}
