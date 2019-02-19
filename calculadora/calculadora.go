package calculadora

import "fmt"

//Calcular: Suma y resta dos numeros ingresados por pantalla
func Calcular() { //con mayuscula la funcion es publica
	var dig1, dig2 int

	fmt.Print("Ingresa primer digito: ")
	fmt.Scanf("%d", &dig1)
	fmt.Print("Ingresa segundo digito: ")
	fmt.Scanf("%d", &dig2)

	suma := suma(dig1, dig2)
	resta := resta(dig1, dig2)

	fmt.Printf("La suma de %d y %d es: %d  \n", dig1, dig2, suma)
	fmt.Printf("La resta de %d y %d es: %d  \n", dig1, dig2, resta)
}

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}
