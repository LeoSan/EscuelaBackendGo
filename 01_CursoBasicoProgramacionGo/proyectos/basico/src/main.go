package main

// Forma manual 
// import (
// pk "proyectos/basico/src/mypackage"
// )

import (
	"fmt"
    mypackage "basico/src/mypackage"
)	

func main(){

	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)
}