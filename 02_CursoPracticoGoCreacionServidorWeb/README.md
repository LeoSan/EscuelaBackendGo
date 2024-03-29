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


## Clase 10: Switch

```
switch val {
case "a":
	fmt.Println("Case A")
case "b":
	fmt.Println("Case B")
case "c":
	fmt.Println("Case C")
default:
	fmt.Println("Invalid input")
}

//ejemplo 
func main() {
	//Imprimimos la orden al usuario
	printMenu()

	//Forma de leer desde consola
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()

	switch operacion {
	case "1":
		fmt.Println("El Resultado de la suma es:", suma())
	case "2":
		fmt.Println("El Resultado de la resta es:", resta())
	case "3":
		fmt.Println("El Resultado de la multiplicación es:", multi())
	case "4":
		fmt.Println("El Resultado de la multiplicación es:", divide())
	default:
		fmt.Println("Invalid input")
	}

	fmt.Println("Adios gracias por usar la calculadora en GO!!")
}
```

## Clase 11-14 : Struct

> Uso de Struct es el equivalente de una clase 

**Nota**
- En cuanto a variables y funciones es importante señalar que la letra inicial de la variables o de la función permite tener visibilidad
- Letra `Mayuscula` es UN METODO PÚBLICO 
- Letra `Minisculas` es un metodo PRIVADO 
- Esto aplica en cuanto a variable y funciones 


```
Las Receiver Function le da la propiedad al struct de tener el método dentro de el, es decir cuando instanciemos a la struct vamos a poder llamar a los métodos dentro de la Receiver Function.
```

```
package main

import "fmt"

//Genero Clase
type task struct {
	name        string
	description string
	complete    bool
}

//Genero metodo
func (t *task) checkTask() {
	t.complete = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de GO",
		description: "Completar el curso de GO en esta semana",
	}
	//fmt.Println(t)
	fmt.Printf("%+v\n", t) //Permite imprimir la estructura de la clase

	t.checkTask()
	t.updateName("Finalizar curso de GO")
	t.updateDescription("Finalizar curso de GO cuanto antes")

	fmt.Printf("%+v\n", t) //Permite imprimir la estructura de la clase

}

```

## Clase 15: Uso de Slices

```
package main

import "fmt"

//Genero Clase

type taskList struct {
	tasks []*task
}

type task struct {
	name        string
	description string
	complete    bool
}

//Genero metodo
func (t *taskList) addTaskList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *task) checkTask() {
	t.complete = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	t1 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de GO",
		description: "Completar el curso de GO en esta semana",
	}
	t2 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de JS",
		description: "Completar este mes ",
	}
	t3 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de Habilidad blanda",
		description: "Completar este mes ",
	}
	t4 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de English",
		description: "Completar este mes ",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.addTaskList(t4)
	fmt.Println(len(lista.tasks))
	lista.deleteTask(1)
	fmt.Println(len(lista.tasks))

}

```

## Clase 16-17 : Ciclo For 

> La mas usada es la de range pero existen 4 tipos de formas 

**Características de range:**

- Es equivalente a foreach en otros lenguajes de programación
- Regresa 2 valores: el índice de la estructura de datos y el valor contenido. En caso de no usar el índice, se puede reemplazar con un guión bajo (_).
- Éste puede iterar sobre valores de tipo string, array, slice, channel ó map.

```
// Sintaxis clásica para definir un ciclo for
for i := 0; i < len(list.tasks); i++ {
    fmt.Println("Index:", i, "task:", list.tasks[i].name)
}
```

## Clase 18: ¿Qué es una interfaz?

> En Go para decir que una clase implementa a una interfaz no es necesario hacerlo de manera explicita como otro lenguajes como c#, Java, etc. En go, se hace de manera implícita. es decir, solo necesita cumplir com las firmas que están en la interfaz.
> Las interfaces en Go simulan el comportamiento de conceptos como clases e interfaces de otros lenguajes, permitiendo también polimorfismo.
> 

```
package main

import "fmt"

func main(){

	figuras := make([]Area, 0)
	figuras = append(figuras, &Cuadrado{lado:4})
	figuras = append(figuras, &Rectangulo{lado:4, base: 6})
	figuras = append(figuras, &Triangulo{base:4, altura:10})

        // Calcular areas
	for _, f := range figuras {
		fmt.Println(f.CalcularArea())
	}

}

type Area interface {
	CalcularArea() int
}

// Cuadrado implementa Area de manera implicita
type Cuadrado struct {
	lado int
}

func (c *Cuadrado) CalcularArea() int {
	return c.lado * c.lado
}

// Rectangulo implementa Area de manera implicita
type Rectangulo struct {
	lado, base int
}

func (r *Rectangulo) CalcularArea() int {
	return r.lado * r.base
}

// Triangulo implementa Area de manera implicita
type Triangulo struct {
	base, altura int
}

func (t *Triangulo) CalcularArea() int {
	return (t.base * t.altura)/2
}
```

## Clase 19-21: ¿Qué es una interfaz?
> Los mapas son muy importantes por una sencilla razón, te permite almacenar información y luego consultarla en tiempo constante.

> Es decir si tú tienes una lista con números únicos por ejemplo, para encontrar un número, debes hacer un for e ir comparando hasta encontrar el correcto.
>Con un mapa tú ya sabes la key específica y la memoria sabe donde está ese pedacito de información, así que irá de una y te regresará el dato que buscabas.

```
package main

import "fmt"

type animal interface {
	makeSound() string
}

type dog struct{}

type cat struct{}

func (d dog) makeSound() string {
	return "Waow waow"
}

func (c cat) makeSound() string {
	return "Miau"
}

func animalMakeSound(a animal) {
	fmt.Println(a.makeSound())
}

func main() {
	cafu := dog{}
	sisa := cat{}

	animalMakeSound(cafu)
	animalMakeSound(sisa)

}

```


## Clase 22: Imprimiendo el contenido de una Página Web usando Interfaces

> usamos lo metodos de GO para poder consumir y excribir desde la web 

`
import (	"io/ioutil"	"net/http" )
`


```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Basic HTTP GET request
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Printf("%s\n", body)

}

```

## Clase 23-24-: Introducción al problema de la Concurrencia

![Explicación](./info/Screenshot_1.png)


## Clase 25-27: Channels

> Un pequeño fact es que este método de comunicación por channels, proviene de un científico llamado Tony Hoare y un concepto que llamó Comunicatin Sequential Processes (CSP) para describir patrones de interacción en sistemas concurrentes

> Básicamente los channels de go parecen tuberías o FIFOS en linux

```
data := <- c // leer de un canal c  
c <- data // enviar data a  canal c 
```

```
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
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

```


## Clase 28-29: Descripción de Servidor Web
> Explica lo que es un cliente-servidor 

**Notas**
- En windows si deseas compilar todos los archivos de ese directororio en windows -> `go run .` en Linux -> `go run *`
- Si te sigue dando error ejecuta este comando `go env -w GO111MODULE=off`


## Clase 30 - 39 : Creamos los archivos correspondientes para crear un servidor 

- Paso 1: Creaos el server [server.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/server.go)
- Paso 2: Creamos los manejadores (handle) [handler.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/handlers.go)
- Paso 3: Creamos los middlerwares  [middlerwares.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/handlers.go)
- Paso 4: Creamos los router  [router.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/router.go)
- Paso 5: Creamos los type -> es nuestro helpers  [types.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/types.go.go)
- Paso 6: Creamos las rutas en el main ->  [main.go](../02_CursoPracticoGoCreacionServidorWeb/proyectos/web-server/main.go.go)
  


# Herramientas GO 
- Convertidor de json - GO -> https://mholt.github.io/json-to-go/
- Convertidor de json - GO -> https://app.quicktype.io/
- Frameeork GI -> https://gin-gonic.com/
- MongoDB -> https://www.mongodb.com/docs/drivers/go/current/
- Standar de GO -> https://github.com/golang-standards/project-layout



## Clase 40 -41 : Go Modules creando nuestro primer módulo en Go

- Lo explicaron muy mal, debo buscar otra información y nutrirme 


# No entendí le pedí ayuda al Chat GTP y pude entender esta clase:
- Paso 1:  Crea un directorio: Lo primero que debes hacer es crear un directorio para el módulo que vas a crear. Este directorio debe tener un nombre que identifique el módulo. Por ejemplo, si estás creando un módulo para manejar fechas, podrías llamar al directorio “fecha”.

- Paso 2:  Crea un archivo go.mod: Dentro del directorio que creaste en el paso anterior, debes crear un archivo llamado go.mod. Este archivo es el que va a contener la información sobre el módulo que estás creando, como el nombre, la versión, las dependencias, entre otros.

- Paso 3: Define el nombre y la versión del módulo: En el archivo go.mod, debes especificar el nombre y la versión del módulo que estás creando. Para hacer esto, escribe el siguiente comando en el archivo:

`module nombre-del-modulo/v1.0.0`

	- Reemplaza `nombre-del-modulo` con el nombre que elegiste para el módulo y `v1.0.0` con la versión inicial del módulo.

- Paso 4: Agrega las dependencias: Si tu módulo depende de otros paquetes o módulos, debes agregarlos al archivo go.mod. Para hacer esto, escribe el siguiente comando en el archivo:

`require (nombre-del-paquete/v1.0.0)`

	- Reemplaza `nombre-del-paquete` con el nombre del paquete que estás agregando y `v1.0.0` con la versión del paquete.

- Paso 5: Crea los archivos del módulo: Ahora debes crear los archivos que van a conformar el módulo que estás creando. Por ejemplo, si estás creando un módulo para manejar fechas, podrías crear un archivo llamado “fecha.go” que contenga el código para manejar fechas.

- Paso 6: Exporta las funciones y estructuras del módulo: Para que otras personas puedan utilizar tu módulo, debes exportar las funciones y estructuras que has creado. Para hacer esto, debes nombrar las funciones y estructuras con una letra mayúscula en la primera letra del nombre. Por ejemplo, si tienes una función llamada “calcularEdad”, deberías nombrarla como “CalcularEdad”.

- Paso 7: Publica el módulo: Una vez que has creado el módulo y has exportado las funciones y estructuras, debes publicarlo en un repositorio para que otras personas puedan utilizarlo. Puedes publicarlo en un repositorio de Git, como GitHub o GitLab, o en un repositorio de Go, como el Registro de Paquetes de Go.

**¡Listo! Con estos pasos, has creado tu propio módulo en Go. Recuerda que para utilizarlo en otros proyectos, debes importarlo en el archivo go.mod de ese proyecto.**

## En Caso de librerias en github Privados

`
Si están trabajando con Go en una empresa, puede que su proyecto tenga dependencias privadas de la misma organización.

Para poder acceder a esos módulos privados tendrán que agregar una variable de entorno a su entorno local, de esa forma go clonará los proyectos utilizando git, en vez de fallar cuando intenta buscarlos a través del proxy de módulos de Go.

Agregar la variable de entorno:
export GOPRIVATE="github.com/[el usuario dueño del módulo]"

# Si usas Go 1.13+ podés ejecutar:
go env -w GOPRIVATE="github.com/[el usuario dueño del módulo]"
Una vez que setees GOPRIVATE, Go va a usar git para clonar los repositorios que coincidan con la variable de entorno. Por defecto, Go clona módulos usando http pero para poder autenticarte con Github y poder clonar repositorios privados vas a necesitar ssh. Para forzar a git a usar ssh en Github, necesitas agregar lo siguiente en tu ~/.gitconfig:

[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
Y voilá! Ya podés sincronizar las dependencias privadas de tu proyecto!

`