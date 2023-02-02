package main

import (
	rt "curso-golang/src/16.Punteros/reto"
	"fmt"
)

func main() {
	myPc := rt.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}