package adivinanumero

import (
	"fmt"
	"math/rand"
)

func Adivina() {
	var numeroMagico = 1 + rand.Intn(10)
	var numeroUsuario int

	fmt.Println("¡Bienvenido a este juego de adivinanzas!")
	fmt.Println("Por favor ingresa un número entre 1 y 10 e intenta adivinar ¡Buena suerte!")
	for {
		fmt.Scanf("%d", &numeroUsuario)

		if numeroUsuario > 10 {
			fmt.Println("El numero ingresado debe ser menor a 10...")
		} else {
			if numeroUsuario != numeroMagico {
				fmt.Println("No has adivinado, intenta nuevamente :D")
			} else {
				fmt.Println("Adivinaste, felicidades :D")
				break
			}
		}
	}
}
