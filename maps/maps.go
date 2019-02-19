package maps

import "fmt"

//GetMap devuelve un map con llave string y valor int
//map estructura llave-valor
func GetMap() map[string]int { //string llave, int valor

	//inicializar mapa forma 1
	mapTest := make(map[string]int)
	mapTest["llave1"] = 1
	mapTest["llave2"] = 10
	fmt.Println(mapTest)

	//inicializar mapa forma 2
	mapTest2 := map[string]int{
		"Juan":  18,
		"Laura": 23,
		"Tatis": 16,
	}

	delete(mapTest, "llave1") //Eliminar llaves
	return mapTest2
}

//GetValorMap Devuelve la edad de determinado nombre
func GetValorMap(name string) int {
	mapTest3 := map[string]int{
		"Juan":  18,
		"Laura": 23,
		"Tatis": 16,
	}
	return mapTest3[name]
}
