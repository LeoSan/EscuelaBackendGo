package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Funciones operacionales

func suma() int {
	var suma int = 0
	counter := 0
	for counter < 1 {
		//Leo Numero
		numero := leeNumero()
		//Valido que sea entero
		sv, err := strconv.Atoi(numero)
		if err != nil {
			fmt.Println("Valida tu teclado solo debes ingresar numero no letras")
			log.Fatal("Sucedio un error tipo: ", err)
			break
		}
		//convierto a entero
		//Sumo
		suma = suma + sv
		//Leo Respuesta si desea continuar si o no
		respuestaScan := leerRespuesta()

		if strings.ToUpper(respuestaScan) != "S" {
			counter++
		}
	}
	return suma
}

func resta() int {
	var resta int = 0
	//Leo Dos Numeros
	numeroA, numeroB := leeNumeroResta()

	svA, err1 := strconv.Atoi(numeroA)
	if err1 != nil {
		fmt.Println("Valida tu teclado solo debes ingresar número NO letras!")
	}

	svB, err2 := strconv.Atoi(numeroB)
	if err2 != nil {
		fmt.Println("Valida tu teclado solo debes ingresar numero NO letras")
	}

	//Calculo la resta
	resta = svA - svB
	return resta
}

func multi() int {
	var multi int = 1
	counter := 0
	for counter < 1 {
		//Leo Numero
		numero := leeNumero()
		//Valido que sea entero
		/*if sv, err := strconv.Atoi(numero); err == nil {
			fmt.Println("entro")
			fmt.Printf("%T, %v\n", sv, sv)
			break
		}*/
		//convierto a entero
		sv, _ := strconv.Atoi(numero)
		//Sumo
		multi = multi * sv
		//Leo Respuesta si desea continuar si o no
		respuestaScan := leerRespuesta()
		if respuestaScan != "S" {
			counter++
		}
	}
	return multi
}

func leerRespuesta() string {

	fmt.Println("Desea continuar S/N : ")
	//Forma de leer desde consola
	scannerResp := bufio.NewScanner(os.Stdin)
	scannerResp.Scan()
	respuestaScan := scannerResp.Text()
	return respuestaScan
}

func leeNumero() string {

	fmt.Println("Ingrese un número: ")
	//Forma de leer desde consola
	scannerSuma := bufio.NewScanner(os.Stdin)
	scannerSuma.Scan()
	numero := scannerSuma.Text()
	return numero
}

func leeNumeroResta() (a, b string) {

	fmt.Println("Ingrese un número A: ")
	//Forma de leer desde consola
	scannerRestaA := bufio.NewScanner(os.Stdin)
	scannerRestaA.Scan()
	numeroA := scannerRestaA.Text()

	fmt.Println("Ingrese un número B: ")
	//Forma de leer desde consola
	scannerRestaB := bufio.NewScanner(os.Stdin)
	scannerRestaB.Scan()
	numeroB := scannerRestaB.Text()
	return numeroA, numeroB
}

func main() {
	//Imprimimos la orden al usuario
	fmt.Println("---------------Menu-----------------")
	fmt.Println("Ingrese la operacion a realizar: ")
	fmt.Println("Ingrese (1) Para realizar una suma(+)")
	fmt.Println("Ingrese (2) Para realizar una resta(-)")
	fmt.Println("Ingrese (3) Para realizar una multiplicacion(*)")
	fmt.Println("---------------Menu-----------------")

	//Forma de leer desde consola
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()

	if operacion == "1" {
		fmt.Println("El Resultado de la suma es:", suma())
	}

	if operacion == "2" {
		fmt.Println("El Resultado de la resta es:", resta())
	}

	if operacion == "3" {
		fmt.Println("El Resultado de la multiplicación es:", multi())
	}

	fmt.Println("Adios gracias por usar la calculadora en GO!!")
}
