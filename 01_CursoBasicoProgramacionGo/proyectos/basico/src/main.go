package main

import (
	"fmt"
	"math"
)

func rArea(base , height int)int {
	return base * height
}

func tArea(higher , less, size int)int {
	return (higher + less) * size / 2
}

func cArea(radio, pi float64)float64 {
	return (radio * radio) * pi
}

func dobleReturn(a1, b1 int)(z, i int){
	var xx int = a1 + b1
	var yy int = a1 * b1
	return xx, yy
}

func main() {

	//Como declarar variables constantes 
	const  pi3 float64 = 3.14
	const  pi2 = 3.14 

	//como declarar variables enteras 
	base2 := 12 
	var altura2 int = 14
	var area2 int 

	fmt.Println("pi:",pi3)
	fmt.Println("pi2:",pi2)
	fmt.Println("Altura:",altura2)
	fmt.Println("area:",area2)
	fmt.Println("base:",base2)
	//Zero values
	//Go asigna vaalores a variables vacías
	var aa int
	var bb float64
	var cc string
	var dd bool

	fmt.Println(aa, bb, cc, dd)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicació:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremetal:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)	

	//Area de un rectangulo, trapecio y de un circulo
	//Rectangulo
	a := 4
	b := 3
	areaRectangulo := a * b
	fmt.Println("Área de un rectángulo:", areaRectangulo)

	//Trapecio
	a = 5
	b = 8
	areaTrapecio := (5 + 8) / 2
	fmt.Println("Área de un trapecio:", areaTrapecio)

	//Circulo
	diametro := 5
	pi4 := math.Pi
	diametroCuadrado := math.Pow(float64(diametro), 2)
	areaCirculo := (pi4 * diametroCuadrado) / 4
	fmt.Println("Área de un circulo:", math.Round(areaCirculo*100)/100)

	base := 80
	height := 40
	fmt.Println("Area del rectangulo: ", rArea(base, height))

	higher := 90
	less := 45
	size := 55
	fmt.Println("Area del trapecio: ", tArea(higher, less, size))

	radio := 45.5
	const pi = 3.14
	fmt.Println("Area del circulo: ", cArea(radio, pi))	

	resultA, resultB :=dobleReturn(12, 12)
	fmt.Printf(" Resultado A = %d , Resultado B : %d ", resultA, resultB)	

	fmt.Printf("\n")	
	fmt.Println("----------------Condicionales--------------------")	
	valorA := 1

	if valorA == 1 {
		fmt.Println("Es 1")	
	}else{
		fmt.Println("No es 1")	
	}

	fmt.Printf("\n")	
	fmt.Println("----------------Switch without a condition--------------------")	

	//Switch without a condition
	value := 50
	switch {
	case value < 0:
		fmt.Println("Value is smaller than zero")
	case value > 100:
		fmt.Println("Value is greater than 100")
	default:
		fmt.Println("Value is between 0 and 100")
	}

	fmt.Printf("\n")	
	fmt.Println("----------------If using a variable is recommended to do this--------------------")	
	switch auxiliar2 := 10; auxiliar2 {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

	fmt.Printf("\n")	
	fmt.Println("----------------//structure of switch--------------------")	
	auxiliar := 10
	switch auxiliar {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

}