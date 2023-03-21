package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Permite
func CheckCreateRequest() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//This function will check if a request is correct or no
			decoder := json.NewDecoder(r.Body)
			var data User
			if err := decoder.Decode(&data); err != nil {
				fmt.Fprintf(w, "error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if data.IsValid() {
				fmt.Printf("%+v\n", data)

				response, err := data.ToJson()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError) //Permite enviar un mensaje de error
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response) //Enviamos el response
				//f(w, r)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Missing or Invalid Fields")
			}

		}
	}
}

// Permite validar el loggin
func Logging() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()                                           //Nos permite saber cuanto tiempo ha ocurrido
			defer func() { log.Println(r.URL.Path, time.Since(start)) }() //defer permite ejcutar una funcion al final //Since te permite cuando tiempo se ejecuta un request
			f(w, r)                                                       //Una  nueva forma de definir funciones son llamadas funciones anonimas
		}
	}
}

// Decide si puede seguir o no simplemente autentica la ruta
// Es un metodo que retorna otro metodo como función
// Middlewer tambien pueden ser handler y puede estar uno dentro de otro
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("checando autenticación!!") //Imprime la hora exacta cuando fue ejecutado
			f(w, r)
		}
	}
}
