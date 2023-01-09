package main

import "fmt"

func main() {

	//Como declarar variables constantes 
	const  pi float64 = 3.14
	const  pi2 = 3.14 

	//como declarar variables enteras 
	base := 12 
	var altura int = 14
	var area int 

	fmt.Println("pi:",pi)
	fmt.Println("pi2:",pi2)
	fmt.Println("Altura:",altura)
	fmt.Println("area:",area)
	fmt.Println("base:",base)
	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)
}