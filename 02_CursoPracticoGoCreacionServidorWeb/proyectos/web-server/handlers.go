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

// Controlador que permitira indicar si estamos en la raiz
// Parametro 1: escribe
// Parametro 2: recibe
// Recuerda es tu request y response
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/ API corriendo")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
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
