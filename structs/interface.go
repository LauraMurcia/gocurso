package structs

//Inter1 es una interfaz de Curso y ProyectoCurricular
type Inter1 interface {
	Inscribir(nombre string)
}

//InterfaceTest Metodo de interfaz
func InterfaceTest() {
	matematicas := Curso{Nombre: "Matematicas", Duracion: "40 horas", Habilidades: []string{"logica", "razonamiento"}}
	ingSistemas := ProyectoCurricular{Nombre: "Ingenieria de sistemas", Cursos: []Curso{matematicas}}

	llamarInscribir(matematicas)
	llamarInscribir(ingSistemas)
}

func llamarInscribir(i1 Inter1) {
	i1.Inscribir("Laura")
}
