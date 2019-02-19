package structs

import "fmt"

// Curso structura de cursos
type Curso struct {
	Nombre      string
	Duracion    string
	Habilidades []string
}

//Inscribir inscribir en un curso
func (c Curso) Inscribir(nombre string) {
	fmt.Printf("la persona %s ha sido inscrita al curso %s \n ", nombre, c.Nombre)
}

//ProyectoCurricular Representa un proyecto curricular
type ProyectoCurricular struct {
	Nombre string
	Cursos []Curso
}

//Inscribir inscribir en un proyecto
func (p ProyectoCurricular) Inscribir(nombre string) {
	fmt.Printf("la persona %s ha sido inscrita al proyecto curricular %s \n ", nombre, p.Nombre)
}

//Estructuras impresion de structs
func Estructuras() {
	//Forma uno de instanciar estructura
	matematicas := Curso{Nombre: "Matematicas", Duracion: "40 horas", Habilidades: []string{"logica", "razonamiento"}}

	//Forma dos de instanciar una estructura
	sociales := new(Curso)
	sociales.Nombre = "Sociales"
	sociales.Duracion = "30 horas"
	sociales.Habilidades = []string{"historia", "geografia"}

	ingSistemas := ProyectoCurricular{Nombre: "Ingenieria de sistemas", Cursos: []Curso{matematicas, *sociales}}

	fmt.Println(matematicas)
	fmt.Println(*sociales)

	fmt.Println(ingSistemas)

	matematicas.Inscribir("Laura")
}
