// el archivo main va a poder leer todo lo que este en el archivo server
package main

import (
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatal("unable to run server", err)
		return err
	}
	log.Print("server running on", s.port)
	return nil
}

// Handle es el nombre de la ruta por ejemplo "/api" asignado a un handler especifico
func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	//Asociacion del handler con la ruta, es decir, el mapa con la llave path asignado al handler
	if !s.router.FindPath(path) {
		//If not path then create a new one
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
	//asi el servidor es capaz de agregar la ruta especifica a un handler especifico
}

// Este metodo permite crear encadenamientos de ejecución de middleware
// ...Middleware -> Agregando ... le permites decir a GO que llegaran varios middlerware o una cantidad de parametros que no sabemos
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	//Agrupo los middlewares y creamos un slice de middleware
	for _, m := range middlewares {
		f = m(f) //m son los multiples middleaware
	}
	return f
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Middleware func(http.HandlerFunc) http.HandlerFunc
