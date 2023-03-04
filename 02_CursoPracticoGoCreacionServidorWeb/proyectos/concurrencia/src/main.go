package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	var i int = 0
	//Creo un canal -> usamos la palabra reservada make
	//Chan indica que es un canal
	// siempre hay que decirle que tipo de string
	canal := make(chan string)

	servidores := []string{
		"https://platzi.com/",
		"https://google.com/",
		"https://facebook.com/",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal) //Asi imprime lo que el canal devuelva
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo transcurrido %s\n", tiempoPaso)

	for i < 2 {

		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

}

func revisarServidor(servidor string, canal chan string) {

	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, "No esta activo =( ")
		canal <- servidor + "No esta activo =( "
	} else {
		fmt.Println(servidor, "Si esta activo =) ")
		canal <- servidor + "Si esta activo =) "
	}
}
