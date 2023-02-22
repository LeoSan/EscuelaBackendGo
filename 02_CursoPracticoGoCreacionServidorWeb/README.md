# Curso Práctico de Go: Creación de un Servidor Web 
> Profesor : Néstor Escoto

## Clase 1-2-3-4-5-6: Mi primer programa en Go

**¿Qué es?**

> Es un Lenguaje compilado (se recopilan los códigos) y estáticamente tipado 
(se debe indicar el tipo de variable o constante para que guarde algún valor en él)

> Se le puede llamar Go/Golang 

**Cuando**
- Anuncio Noviembre 2009
- Primera version 2012

**Creado**
- Por Google​ y sus diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson
- 

**Caracteristicas**
- Maneja procesos pesados, es potente, pero amigable.
- Se utiliza Go/Goland para nombrarlo.
- Los programadores de este lenguaje se hacen llamar gophers.
- Es veloz
- Tiene alto rendimiento para tareas pesadas
- Maneja soporte nativo por concurrencia
- Un Gopher puede ganar $74k al año
- Facilita ajustar sintaxis de forma nativa
- Comunidad receptiva, contribuye y apoya.
- Tenemos un compilar Online `https://go.dev/tour/moretypes/7`

**¿Dónde se usa?**
- Mercado Libre
- Twich
- Twitter
- Uber
- Docker y Kubernetes


## Clase 7: Leer inputs desde la consola 

> Para leer desde consola usamos el paquete `bufio` debemos invocarlo al y luego usarlo de esta manera

- Para este ejemplo cree una función para simplicar el uso 
```
import (
	"bufio"
	"fmt"
	"os"
)

func leeNumero() string {

	fmt.Println("Ingrese un número: ")
	//Forma de leer desde consola
	scannerSuma := bufio.NewScanner(os.Stdin)
	scannerSuma.Scan()
	numero := scannerSuma.Text()
	return numero
}

```

- Paso 1: importamos `bufio` y `os`
- Paso 2: Declaramos una variables que al mismo tiempo permite obtener lo que tecleas en consola `scannerSuma := bufio.NewScanner(os.Stdin)`
- Paso 3: Dicha variable lee y almacena lo que se teclea `scannerSuma.Scan()`
- Paso 4: extraemos de la variable scan indicandole que es un text `scannerSuma.Text()`
- Paso 5: Listo podemos mostrar lo que tecleamos en consola con un `fmt.Println("tu tecleaste", numero)`

## Clase 8: Manejo de errores y uso de If

> Sigue siendo confuso este tema pero si deseamos manejar un error al convertir valores podemos usar el paquete `strconv` y declarando una variable err podemos manejar el resultado en multiples variables ejemplo: 

```
        numero := leeNumero()
		//Valido que sea entero
		sv, err := strconv.Atoi(numero) // Aqui declaro dos variables si esta bien lo guarda en SV si esta mal lo guarda en ERR

        //Valido lo que viene en este caso si esta mal llega con nil
		if err != nil {
			fmt.Println("Valida tu teclado solo debes ingresar numero no letras")
		}
```

> Podemos manejarlo con una función general 

```
//Nota para manejar mensajes de este tipo debemos importar ""log""

func validaError(err error){
    if err != nil {
        log.Fatal("Sucedio un error tipo: ", error)
    }
}
```

## Clase 9: Mayusculas y Minusculas

> Para este caso usamos el paquete `strings` y los metodos para Mayuasculas -> `strings.ToUpper(respuestaScan)` y para Minusculas -> `strings.ToLower(respuestaScan)`
```
		//Mayusculas 
        if strings.ToUpper(respuestaScan) != "S" {
			counter++
		}
		//Minusculas  
        if strings.ToLower(respuestaScan) != "s" {
			counter++
		}
```        
