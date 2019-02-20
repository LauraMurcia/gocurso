package holamundo

import (
	"errors"
	"fmt"
	"strings"
)

const helloWorld string = "Hola  %s %s, bienvenido"

func holaMundo() {
	name := GetName()
	lastname := "<Modificar con mi apellido>"
	number := sum(2, 3)
	a, b, c := getVariables()
	f32, f64 := getFloat()

	fmt.Printf(helloWorld, name, lastname)
	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Hola"[0])         //Devuelve ascii de H
	fmt.Println(string("Hola"[0])) //Devuelve el valor H
	fmt.Println("cantidad de caracteres: ", len("Hola"))
	getArray()
	getSlice()
	esPar()
	forTest()
	stringTest()
	switchTest()
}

func GetName() string {
	var name string

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)

	return name
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func getVariables() (int, int32, int64) {
	return 1, 2, 3
}

//Suma2 suma con errors
func Suma2(a interface{}, b interface{}) (int, error) {

	switch a.(type) { //a t le asigno tipo de dato de a
	case string:
		return 0, errors.New("El valor A es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}

func sum(a int, b int) int {
	return a + b
}

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "array"
	arr1[1] = "array2"

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func getSlice() { //No se especifica tamanio
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

func esPar() {
	var number = 0
	fmt.Println("Ingresa un numero del 1 al 10")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("El numero ingresado es par")
	} else {
		fmt.Println("El numero ingresado es impar")
	}
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("For 1 condicion: ", c)
	}

	//Simulacion while
	s := 1000
	for {
		s -= 10
		if s == 0 {
			fmt.Println("Termina la operacion")
			break
		}
	}
}

func stringTest() {
	var texto = "Hola Mundo, Hola platzi, Hola Go"

	fmt.Println(strings.ToUpper(texto))                      //Mayuscula
	fmt.Println(strings.ToLower(texto))                      //Minuscula
	fmt.Println(strings.Replace(texto, "Hola", "Hello", -1)) // -1 para que los reemplace todos
	fmt.Println(strings.Split(texto, ","))
}

func switchTest() {

	var mes = 0
	fmt.Println("Ingrese el número del mes: ")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	default:
		fmt.Println("Mes no válido")
	}

	//switch sin expresion, evalua la condicion en cada case
	switch {
	case mes%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}
}
