# Curso de Go Intermedio: Programación Orientada a Objetos y Concurrencia
> Profesor : Néstor Escoto

## Clase 1: Características esenciales de Go


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
- Lenguaje compilado
- Potente librería estándar
- Manejo de concurrencia a través de GoRoutines y Channels.
- Diseñado por Google, Ken Thompson (UNIX) parte del equipo de diseño
- Muy popular en aplicaciones CLI y Backend
- Docker, Kubernetes y Terraform están escritos en Go.
- Muy utilizado para escribir malware.
- Según StackOverflow, es el tercer lenguaje mejor pagado a nivel global.
- Según StackOverflow, es la quinta tecnología mas amada por los desarrolladores y la tercera mas deseada para trabajar.
- Se maneja las variables de manera por `referenia` para evitar el duplicado de variables en memoria 



**¿Dónde se usa?**
- Mercado Libre
- Twich
- Twitter
- Uber
- Docker y Kubernetes


## Clase 2: Qué aprenderás y qué necesitas saber

**Conocimientos Requeridos**

- Declaración de variables.
- Condicionales.
- Sintaxis básica.
- Declaración de GoRoutines y Channels.
- Slices y maps.
- Apuntadores.

**¿Qué terminarás sabiendo?**
- ¿Es Go orientado a objetos?
- Cómo aplicar los conceptos de POO en Golang.
- Crear Test Unitarios en Go para aplicar TDD.
- Calcular el Code Coverage de mi proyecto.
- Análisis del Profiling en tus programas.
- Cómo multiplexar mensajes de canales.
- Técnicas para la creación de Worker Pools concurrentes.
- Crear Test Unitarios en Go para aplicar TDD.
- Crear un Job Queue concurrente

## Clase 3: Repaso general: variables, condicionales, slices y map


**Ejmeplo código**
```
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Capturando valor y error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

	// Validando errores de manera explicita“.
	// Nil = valor nulo
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa: estructura de clave valor, con Make especificar un map que mapee llaves de tipo string a valores de tipo entero
	m := make(map[string]int)
	//Key= string, 6= int
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice: Estructura como un array
	s := []int{1, 2, 3}
	// Con el for recorremos el slice
	for index, value := range s {
		//index: valor en memoria que estamos accediendo
		fmt.Println(index)
		//value: valor almacenado en el slice
		fmt.Println(value)
	}
	// Agregar un valor nuevo al final del slice
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
}
```

**Ejmeplo código**
```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// Handling errors
	myValue, err := strconv.ParseInt("Nan", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", myValue)
	}

	// Map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	// Slice
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println()

	// Append slice
	s = append(s, 6)

	// Range
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println()

	// Delete from slice
	s = append(s[:2], s[3:]...)

	// Range
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Bob", 20}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

}
```

# Clase 4: Repaso general: GoRoutines y apuntadores

- Temas para dominar -> https://divan.dev/posts/go_concurrency_visualize/

**Ejmeplo código**
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //Creamos un canal para monitorear las Goroutines
	go doSomething(c)   //Llamamos a la función con go para generar un Goroutine
	<-c                 //main esperara a que este canal reciba el mensaje

	g := 25
	fmt.Println(g)
	h := &g         //Apuntador a g
	fmt.Println(h)  //Dirección de memoria donde esta almacenada g
	fmt.Println(*h) //Accedemos al valor de h que a su vez almacena g
}

func doSomething(c chan int) { //Recibimos un canal de tipo int
	time.Sleep(3 * time.Second) //Poner a dormir
	fmt.Println("Done")
	c <- 1 // le envia el valor de 1
}
```

## Clase 5: ¿Es Go orientado a objetos?

**¿Es Go Orientado a objetos?**

- POO se ha convertido en uno de los paradigmas de programación predominante en la industria.
- POO puede llegar a ser muy riguroso, pero a cambio permite una alta reutiliuzación de código y la aplicación de un sinnúmero de patrones de diseño.
- Go puede alcanzar la aplicación de los conceptos de POO, pero de una forma diferente a lenguajes como Java y Python.

**GO resalta en lo siguiente**

1.- Encapsulación
	- a. Estado(“campos”)
	- b. Comportamiento (“métodos”)
	- c. Exportado & No exportado; visible & no visible
2.- Reusabilidad
	- a. herencia (“tipos embebidos”)

3.- Polimorfismo
	- a. interfaces
4.- Overriding
	- a. "promoción"
	- OOP Tradicional
5.- Clases
	- a. Estructura de dato describiendo un tipo de objeto
	- b. Puedes crear “instancias”/ “objetos” de la clase / prototipo
	- c. Las clases almacenan ambos:
		- i. estado / datos / campos
		- ii. comportamiento / métodos
	- d. público / privado
6.- Herencia
	- En Go:
	- No creas clases, creas un TIPO
	- No creas instancias, creas un VALOR de un TIPO
	- Tipos definidos por el usuario
	- Podemos declarar un nuevo tipo

**Tipos Nombrados vs Tipos Anónimos**

- Tipos anónimos son indeterminados. Aún no han sido declarados como un tipo. El compilador tiene flexibilidad con tipos anónimos. Puedes asignar un tipo anónimo a una variable declarada de cierto tipo. Si la asignación puede ocurrir, el compilador hará el trabajo de determinar el tipo; el compilador hará una conversión implícita. No puedes asignar un tipo nombrado a un tipo de diferente nombre.
- Alineamiento arquitectónico y más

**Convención: organiza tus campos lógicamente. **
- Legibilidad y claridad ganan en rendimiento como punto crítico. 
- Go será de buen rendimiento. 
- Ve primero por legibilidad. 
- Sin embargo, si estás en una situación donde necesitas darle prioridad al rendimiento: agrega los campos del más grande al de menor tamaño, por ejemplo: int 64, int32, float32, bool

## Clase 6: Structs vs. clases



**Ejmeplo código**
```
//Es la forma de generar una clase y  una instancia en go 
// Usamos Struct NO class
// usamos referncia usando * & como una instancia 

package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Println("%v", e)

	e.id = 1
	e.name = "Name"
	fmt.Println("%v", e)
}
```

**Concepto**
- El objetivo de una clase es definir una serie de propiedades y métodos que un objeto puede usar para
alcanzar diferentes objetivos.
- Go utiliza Structs para generar `nuevos tipos` de datos que se pueden utilizar para cumplir tareas en
específico.

**Notas**
- Es la forma de generar una clase y  una instancia en go 
- Usamos Struct NO class
- usamos referncia usando * & como una instancia 


## Clase 7: Métodos y funciones

**Concepto**
- podemos generar metodos getter y setter parecido en java 
- solo necesitamos seguir el patrón colocamos la palabra `func` seguido del nombre de la función seguido definiendo antes que tipo de valor devolvera y que tipo de valor recibirá 
- (`*`) notese el astericos hace una copia de referencia del struct 
- Algunos lenguajes de programación implementan la filosofía que TODO debe ser un objeto, sin embargo, no siempre es algo aplicable, por ejemplo, en una librería con utilidades que no pertenecen a un dominio específico
```
func (p *Employee ) setId(id int){
	p.id = id
}
```

**Ejemplo código**

```

package main
import (
	"fmt"
	// "strconv"
)

type Employee struct{

	id int
	name string
}

func main(){
	p := new(Employee)
	 
	p.setId(1)
	p.setName("Javier")
	p.getDatos()

}

func (p *Employee ) setId(id int){
	p.id = id
}

func (p Employee ) getId() int{
	return p.id 
}

func (p *Employee ) setName(name string){
	p.name = name
}

func (p Employee ) getName() string{
	return p.name
}

func (p Employee) getDatos(){
	fmt.Println(p)
}
```

## Clase 8: Constructores

**Concepto**
> Go no es 100% orientado a objectos, para hacernos una idea si Java es 95 % orientado a objetos (Recuerden que los primitivos de Java no son objectos) go es como el 50%, talvez un poco más un poco menos.


**Creando una clase**
> Para crear una clase en go utlizamos un struct y tambien debemos utilizar la palabra reservada type al comienzo de la declaración.
```
// Esta es nuestra clase con dos atributos
type Employee struct {
	id   int
	name string
}
```


**Creando un constructor**
- Pensemos en un constructor como un representante de un gran artista el cual vela por los intereses del artista que está representando.
- De igual forma el constructor de la clase vela por los intereses de la clase a la que representa, por ejemplo el constructor determina cómo es que se debe instanciar a una clase.
- Aunque el profesor mostró varios ejemplos para crear un constructor a mi me gustaron estos dos.

```
//Forma No1: Crear directamente el objeto donde se va a utilizar .
func main() {

	e1 := Employee{
		id:   14,
		name: “Esteban”,
	}
	fmt.Printf("%v\n", e1)
}

//Forma No2: Crear una función que reciba por parametro los valores que tendrá nuestro objecto , asigne los valores a sus respectivos atributos y luego devuelva una referencia en memoria del objeto creado.

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,  //Asignando valores.
		name: name,
	}
}

//Nuestra función principal
func main() {
	e2:= NewEmployee(23, "Miguel") //Instanciando nuestro objecto.
	fmt.Printf("%v\n", *e2)

}

```

**Creando métodos.**
- En simples palabras un método determina el comportamiento que tendrá un objeto. 
- Utilicemos los famosisimos métodos getters y setters

```
// Este es nuestro método set, y tenemos que indicarle que pertenece a nuestra clase Employee para ello le pasamos un puntero * que apunte a esa clase  ​para luego poder tener acceso al atributo y asignarle el valor que nos pasen a través de este método setId.
func (x *Employee) setId(i int) {
	x.id = i
}

// Este es nuestro método get, igual que el anterior pasamos un puntero * pero en este caso vamos a retornar un valor int por ello lo indicamos en la declaracion del metodo.
func (x *Employee) getId() int {
	return x.id
}

//Nuestra función principal
func main() {

	e2:= NewEmployee(23, "Miguel")
	fmt.Printf("%v\n", *e2)
	idE2 := e2.getId() // recuperamos el atributo id del objecto e2
	e2.setId(14) // Cambiamos el atributo id del objecto e2
	fmt.Printf("%d\n", idE2)
	fmt.Printf("%v\n", *e2)
}

```

## Clase 9: Herencia

**Concepto**
- Go no permite la herencia, go utiliza la `composicion`.
- La composicion, a diferencia de la herencia, no es una clase hija, sino que contiene los metodos de las clases indicadas.
- usamos la palabra reservada `type` para emular herencia 
- Cuando se tienen `receiver functions` asociadas a un struct y se forma otro struct a partir de este, los métodos también funcionan con el nuevo struct.
- Y aún contando con esa herencia de métodos puedo reescribir la receiver function en el nuevo struct solo creando una función con el mismo nombre y recibiendo el struct. Así tendría polimorfismo.

```
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func getMessage(p Person) {
	fmt.Printf("%s with age %d", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
}
```

## Clase 10: interfaces 

**Concepto**
- En cambio en Go, es suficiente definir los métodos para implementar la interfaz y la implementación es implícita.
- En el polimorfismo se utiliza una interfaz (o una clase base) para determinar en tiempo de ejecución el método a utilizar, 
- Se permite acceder por medio de la interfaz en vez de hacerlo a través de un objeto en particular, porque en este último caso el método se determina en tiempo de compilación (esto no es polimorfismo).
- GO no implementa inetrfaces de manera explicita 

```
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Inheritance tipo anonimo
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemproaryEmployee struct {
	Person
	Employee
	taxRate float64
}

type PrintInfo interface {
	getMessage() string
}

// Composicion sobre herencia
func (ft FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a full time employee", ft.name, ft.age)
}

func (te TemproaryEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a temprary employee", te.name, te.age)
}

func GetMessage(pi PrintInfo) {
	fmt.Println(pi.getMessage())
}

func main() {
	f := FullTimeEmployee{}
	f.name = "John"
	f.age = 30
	f.id = 1
	f.endDate = "2019-01-01"
	fmt.Printf("%+v\n", f)

	t := TemproaryEmployee{}
	t.name = "Mary"
	t.age = 25
	t.id = 2
	t.taxRate = 0.25
	fmt.Printf("%+v\n", t)

	//Lo que nos indica que podemos usar el mimos metodo pero implementando diferentes comportamiento 
	GetMessage(f) //-> f empleado tiempo completo 
	GetMessage(t) //-> t empleado temporal 
	//Notese que se usa el mismo metodo pero para diferentes comportamientos  logramos el comportamiento polimorfismo
}
```