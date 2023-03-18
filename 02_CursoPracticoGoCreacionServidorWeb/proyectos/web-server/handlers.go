// el archivo main va a poder leer todo lo que este en el archivo router
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 20)
	fmt.Fprintf(w, "Ruta Home, mostrando un mensaje simple, aun no es una pagina html!")
}

// Handler para manejar la ruta principal
// Parametros: escritor w del tipo http.ResponseWriter y objeto request
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	//Impresion en el navegador
	//parametros: escritor-objeto encargado de responder al cliente
	//y mensaje escrito a travez del escritor
	fmt.Fprintf(w, "/ API corriendo")
}

// Creamos metodo que permite ejeuctar ordenes tipo POST
// usamos librerias NewDecoder permite decodificar un Json
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body) //Body viene la data Json enviada por request
	var metadata MetaData

	//Valido si el json viene con algun error
	if err := decoder.Decode(&metadata); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Request Payload: %v\n", metadata)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything seems to be fine!")
}

type MetaData interface{}
