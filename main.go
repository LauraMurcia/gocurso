package main

import (
	"fmt"
	"time"
	/*
		"github.com/LauraMurcia/gocurso/adivinanumero"
		"github.com/LauraMurcia/gocurso/calculadora"
		"github.com/LauraMurcia/gocurso/holamundo"
		"github.com/LauraMurcia/gocurso/maps"
		"github.com/LauraMurcia/gocurso/structs"*/)

func main() {
	/*calculadora.Calcular()
	adivinanumero.Adivina()

	nombre := holamundo.GetName()
	fmt.Println("Nombre ingresado: ", nombre)

	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetValorMap("Laura"))

	structs.InterfaceTest()

	number, err := holamundo.Suma2(50, 20)

	if err != nil {
		panic(err) //Sirve para detener la ejecucion
	}
	fmt.Println(number)*/

	//pointerTest()
	go forGo(500)
	go forGo(400)
	time.Sleep(10000 * time.Millisecond) //Esperar para que no se cancele operacion
}

func pointerTest() {
	a := 100
	var b *int //* indica que se recibe espacio en memoria
	b = &a     //con &asigamos direccion en memoria
	fmt.Println(a, *b)
	fmt.Println(&a, b) // con *variable pasamos el valor
	*b = 10            //para modificar el vaor de b se usa *b
	fmt.Println("DEspues de modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func helloGo(index int) {
	fmt.Println("Hola soy un print en la Goroutine # ", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
