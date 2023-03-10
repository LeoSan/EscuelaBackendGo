# Curso Básico de Programación en Go 
> Profesor : Osmandi Gómez

## Clase 1: ¿Qué es, por qué y quienes utilizan Go?

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
- Es fuertemente tipado 


**¿Dónde se usa?**
- Mercado Libre
- Twich
- Twitter
- Uber
- Docker y Kubernetes


## Clase 2: Instalar Go en Linux

- Paso 1: Debemos descargar go accediendo a este acceso -> [Descarga Go](https://go.dev/dl/)
- Paso 2: 
- Paso 3: Debemos crear una variable de usuario -> `export GOPATH=/home/usuario/go` -> Entorno donde programas 
- Paso 4: Debemos crear una variable de usuario -> `export GOBIN=$GOPATH/bin`       -> Entorno del codigo Binario 
- Paso 5: Debemos crear una variable de usuario -> `export GOROOT=/usr/local/go`    -> Entorno donde esta el ejecutador 
- Paso 6: Debemos crear una variable de usuario -> `export PATH=$PATH:$GOBIN:$GOROOT/bin`    -> 
- Paso 7: source ~/.bashrc ó . . bashrc 
- Paso 8: podemos ejecutar el comando para validar la instalación `go version`

> PD: 
- mkdir -p $HOME/go/{bin,pkg,src} -> Permite generar el entorno 
- directorio *bin* -> Guarda los ejecutables 
- directorio *pkg* -> Guarda codigo dependencia 
- directorio *src* -> Aqui es donde programas  
- sudo apt-get --only-upgrade install golang -> Instalación mas rapida.

## Clase 3: Instalar Go en Windows

> Pues es sencillo solo descargando el instalador de la pagina oficial -> [Descarga Go](https://go.dev/dl/)
- Paso 1: Instalas con el wizard 
- Paso 2: Compruebo con el comando `go version`
- Paso 3: en caso de error valida las variables de entorno ya el wizard las instala pero si hay error valida o reinicia el equipo 

## Clase 4 - 5 - 6:  Nuestras primeras líneas de código con Go

**Como compilar y ejecutar GO**
- Paso 1: recuerda que debes tener una estructura para trabajar src pkg bin, nuestro caso iniciaremos en src. 
- Paso 2: nuestro proyecto debe quedar algo así nombreProyecto/src/main.go al igual que JAVA este también maneja un main como buena practica. 
- Paso 3: Luego de escribir el código base o inicial debemos ir a consola y ejecutar el siguiente comando `go build ./src/main.go` este compila y genera el exe 
- Paso 4: Luego de validar que se creo el main.exe podemos ejecutarlo con el siguientes comando `./main`
- Paso 5: Vemos en consola en este caso el hola mundo. 

**Otra Forma**
- Podemos ejecutar el código sin compilar pero no es recomendable pero es mas rapido
- Este compila, ejecuta y elimina el código. 
- `go run .\src\main.go`

> Checa esta imagen de referencia. 
![Compilacion_ejecucion_go](./info/Compilacion_ejecucion_go.png)

## Clase 7: Como podemos declarar variables  

**Constantes**
```
	//Como declarar variables constantes 
	const  pi float64 = 3.14
	const  pi2 = 3.14 

``` 

**Enteros**
```
	//como declarar variables enteras 
	base := 12 
	var altura int = 14
	var area int  
``` 

**Zero values**
```
	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

    //Salida -> 0 0 ' ' false
``` 
**Notas:**
- En go cada variable declarada debe usarse esto es para evitar variables que no se usen, todo con el proposito de ahorrar memoria 
- Go asigna valores a variables vacías

```
func main() {
	primeraForma := 1 //Bastante Parecida a Python
    	var segundaForma = 2 //Como en JavaScript
	var terceraForma int16 = 3 // Parecido a C++
}
```

**Zero Values**
>Algo que se debe tomar en cuenta en GO es que si no especificas el valor de tu constante o variable a diferencia de otros lenguajes que le asignan un valor null, GO le dara los valores por Default.

**Valores por Default:**
- int y sus derivados(int8,int16,uint,etc.) : 0
- float32 y float64 : 0
- string : Un string vacio: " "
- bool : False

## Clase 8: Operadores aritméticos
> Las operaciones se hacen clásicamente. + para suma, - para resta, * para multiplicación, / para división y % para módulo, También a tener en cuenta, que en python, una división, sea cuál sea el resultado, lo convierte en un flotante, en Go no.

**Ejemplo**
```

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicació:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremetal:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)	

	//Area de un rectangulo, trapecio y de un circulo
	//Rectangulo
	a := 4
	b := 3
	areaRectangulo := a * b
	fmt.Println("Área de un rectángulo:", areaRectangulo)

	//Trapecio
	a = 5
	b = 8
	areaTrapecio := (5 + 8) / 2
	fmt.Println("Área de un trapecio:", areaTrapecio)

	//Circulo
	diametro := 5
	pi := math.Pi
	diametroCuadrado := math.Pow(float64(diametro), 2)
	areaCirculo := (pi * diametroCuadrado) / 4
	fmt.Println("Área de un circulo:", math.Round(areaCirculo*100)/100)
```

## Clase 9: Tipos de datos primitivos

> Al codificar en Go podemos especificar el tipo de dato, permitiéndonos ganar gran preformas en nuestro código

**Números Enteros**
> int Cuando no se declara el tamaño tomara la referencia del OS (Sistema operativo) (32 o 64 bits)

- int8 8bits ⇒ -128 a 127
- int16 16bits ⇒ -2^15 a 2^15-1
- int32 32bits ⇒ -2^31 a 2^31-1
- int64 64bits ⇒ -2^63 a 2^63-1

**Optimizar memoria cuando sabemos que el dato siempre va ser positivo**

- uint ⇒ Depende del OS (32 o 64 bits)
- uint8 ⇒ 8bits = 0 a 127
- uint16 ⇒ 16bits = 0 a 2^15-1
- uint32 ⇒ 32bits = 0 a 2^31-1
- uint64 ⇒ 64bits = 0 a 2^63-1

**Números decimales**
- float32 ⇒ 32 bits = +/- 1.18e^-38 +/- -3.4e^38
- float64 ⇒ 64 bits = +/- 2.23e^-308 +/- -1.8e^308

**Textos y booleanos**
A diferencia de otros lenguajes de programación donde para definir una variable de tipo string es permitido utilizar “”, ‘’ o ```` en Go solo podemos utilizar las comillas dobles ""

**boolean ⇒ Trueo False**

**Números complejos**
- Complex64 ⇒ Real e Imaginario float32
- Complex128 ⇒ Real e Imaginario float64

**Ejemplo //Data Types**

```
	var shortInt int8 = 3
	var longInt int64 = 2313212113234256876
	var shortFloat float32 = 3.1
	var longFloat float64 = 3.153465212456432145668723312
	var text string = "string"
	var boolean bool = true
	var complex complex128 = 10 + 8i
```	

>PD: Marca la diferencia si definimos el tipo de dato ya que esto permite tener un código mucho mas rapido y eficiente. 

## Clase 10: Paquete fmt: algo más que imprimir en consola
> El paquete fmt es el que se encarga de administrar los inputs y outputs de la terminal.
>

**Tipos de Print:**
- Println: Es un print normal con un salto de linea al final. Ejemplo:
```
fmt.Println("Hola Mundo")
```

- Printf: Es un print al cual le puedes especificar el tipo de objeto que le vas a dar. Ejemplo:
```
fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // Se agrega %v cuando no sabes el tipo de dato
```
- Sprintf: No imprime nada en consola, simplemente lo guarda como un String. Ejemplo de uso:
```
var message string = fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
fmt.Println(message)
```

- Printf: Con este paquete podemos imprimir en consola el tipo de dato de variables o constantes. Ejemplo de uso:
```
package main

import "fmt"

func main() {

	const nombre string = "UltiRequiem"

	fmt.Printf("La variable 'nombre' es de tipo : %T\n", nombre)
}
```
>PD-> %s = significa que el valor sera un String - %d = significa que el valor sera un entero  int - %T = identificar el tipo de datof

**Mas Formatos**
- https://gobyexample.com/string-formatting  
- https://pkg.go.dev/fmt

## Clse 11 : Uso de funciones

>Notas: 
- La palabra reservada es `func`
- Es importante que se usa camelCase por lo que debes siempre poner mayuscula
- Puedes reciclar la declaración de variables en la función `tArea(higher , less, size int)`
- Debemos especificar el tipo de valor que va retornar `tArea(higher , less, size int) int`
- Podemos retornar dos parametros ejemplo  `doubleReturn(higher , less, size int) (a, b int)`
- Como podemos recibirlos ejemplo  `value1, value2 := doubleReturn(2)`
- Usando piso ayuda a descartar cuando desea validar cual retornar `value1, _ := doubleReturn(2)`

```
package main

import "fmt"

func rArea(base , height int)int {
	return base * height
}

func tArea(higher , less, size int)int {
	return (higher + less) * size / 2
}

func cArea(radio, pi float64)float64 {
	return (radio * radio) * pi
}

func dobleReturn(a1, b1 int)(z, i int){
	var xx int = a1 + b1
	var yy int = a1 * b1
	return xx, yy
}

funcmain() {
	base := 80
	height := 40
	fmt.Println("Area del rectangulo: ", rArea(base, height))

	higher := 90
	less := 45
	size := 55
	fmt.Println("Area del trapecio: ", tArea(higher, less, size))

	radio := 45.5
	const pi = 3.14
	fmt.Println("Area del circulo: ", cArea(radio, pi))

	resultA, resultB :=dobleReturn(12, 12)
	fmt.Printf(" Resultado A = %d , Resultado B : %d ", resultA, resultB)		
}
```

## Clase 12 : Go doc: La forma de ver documentación

> Podemos usar la documentación de goland `https://pkg.go.dev/`

## Clase 13 : El poder de los ciclos en Golang: for, for while y for forever

> En Golang solo existe un ciclo for, pero este tiene tipos de variantes: 

- For condicional

```
for i := 0; i<=10; i++ {
	fmt.Println(i)
}
```

- For while
```
counter := 0

for counter < 10 {
	fmt.Println(counter)
	counter++
}
```


- For forever
```
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 12 {
			break
		}
	}
```

- For Range
``` 
	arreglo := [8]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arreglo:", arreglo)

	fmt.Println("Primer ejemplo")
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}

	fmt.Println("Segundo ejemplo")
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}

	fmt.Println("Tercer ejemplo")
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}
```	
- For con funciones
```
for i := preFor(); condicion(i); i = postFor(i) {
		fmt.Printf("Valor de i: %d", i)
		if i == 7 {
			fmt.Printf(" así que saldremos del ciclo...\n")
			break /// este ejemplo es para usar el break
		}
		fmt.Printf("\n")
	}
```
- For con goto tags
```
var i int
CICLO:
	fmt.Println("estamos fuera del for")
	for i < 10 {
		if i == 6 {
			i = i + 3
			fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
			goto CICLO2
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
CICLO2:
	fmt.Printf("ciclo 2 Valor de i: %d\n", i)
	if i == 9 {
		fmt.Printf("Valor de i: %d\n", i)
		i = i + 3
		fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
		goto CICLO
	}
	fmt.Printf("terminamos\n")

```
## Clase 14 : Operadores lógicos y de comparación
> Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE.

>Nota: 
- Como ves no se usan () parentesis para las condicionales ejemplo `if valor == 1 {}`


```
valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.
```

- Convertir texto a entero -> https://pkg.go.dev/strconv

## Clase 16 : Múltiple condiciones anidadas con Switch

> Como bien sabemos es otra forma de manejar condicionales pero en este caso es mas ordenado y facil de leer en GOland podemos crear  tres tipos de Switch >

- //structure of switch
- //If using a variable is recommended to do this
- //Switch without a condition

```
package main

import "fmt"

func main() {
	//SWITCH

	//structure of switch
	auxiliar := 10
	switch auxiliar {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

	//If using a variable is recommended to do this

	switch auxiliar2 := 10; auxiliar2 {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

	//Switch without a condition
	value := 50
	switch {
	case value < 0:
		fmt.Println("Value is smaller than zero")
	case value > 100:
		fmt.Println("Value is greater than 100")
	default:
		fmt.Println("Value is between 0 and 100")
	}

}
```

## Clase 17 : El uso de los keywords defer, break y continue

**keywords** `defer`
- Es el keyword que va ejecutar la ultima función antes que todo muera. 
- Como usarlo como practica, para cerrar conexiones de bases de datos ó cerrar un archivo leido
- Solo se debe usar un defer por función 
```
package main

import "fmt"

func main() {
   fmt.Println("Inicio")
      for i := 0; i < 5; i++ {
      defer fmt.Println(i)
   }
   fmt.Println("Fin")	
}
```

**keywords** `continue y break`
- Es usado en los ciclos con la habilidad de detener `break` o romper un ciclo o de continuar un ciloc `continue`

```
import"fmt"

func main(){
	// Continue y break
	for i := 0; i < 10; i++{
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
	}
} 
``` 


## Clase 18 : Arrays y Slices

> Existen dos tipos para manejar vectores Arrays y slices 

**Arrays**
- No podemos agregar otro elementos 
- Si podemos editar los valores de los elementos
- podemos usar el método `len()` indicando tamaño actual del array
- Podemos usar el método `cap()` indicando la capacidad máxima del array 


```
var array[4] int
array[0] = 1
array[1] = 2
fmt.Println(array, len(array), cap(array))
```

**Slices**
- Similar al arrays pero no se le indica la capacidad total 
- podemos usar el método `len()` indicando tamaño actual del array
- Podemos usar el método `cap()` indicando la capacidad máxima del array 
- La ventaja de slice es que puede recibir indice y este nos puede imprimir por rango los valores del elemento
- Por rango el primero es inclusivo y el segundo exclusivo 
- Los Slice son mutables si podemos agregar elementos  
- uso de propt`...` para agregar nuevos elementos en pocas palabras descompone 

```
slice := []int {0,12,3,4,5,6}
//Métodos en el Slice 
fmt.Println(slice, len(slice), cap(slice))
fmt.Println(slice[0])
fmt.Println(slice[:3])
fmt.Println(slice[2:4])
fmt.Println(slice[4:])

//Agrego 
slice = append(slice, 7)
fmt.Println(slice)

//Agrego lista 
newSlice:= []int{8,9,10}
slide = append(slice, newSlice...)

```
## Clase 19 : Recorrido de Slices con Range

```
package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("anna"))
}

// Pasar Mayusculas 

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Ama"

	isPalindrome(text)
}

func isPalindrome(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
```

## Clase 20: Llave valor con Maps
> Los maps/hash son más eficientes según el caso de uso, si conoces el KEY siempre será más eficiente usar un MAP, debido a que el tiempo de acceso es EN PROMEDIO constante (O(1)) para todos los keys independientemente del tamaño del MAP.

- Map son diccionarios en GO 
- Se usa para manejar valores constantes 
- Hay que tener cuidado con el recorrido ya que es concurrente se puede ejecutar de forma aleatoria 
- Si solo necesitas iterar sobre un conjunto de elementos, siempre es mejor alternativa usar un Array.
- Usa concurrencia de manera nativa 
- Arrays y Hashes, son bastante similares en memoria, los hashes son en resumen un hack para crear arreglos en los que podemos definir un KEY distinto a un índice entero.

- Este es un muy buen artículo para conocer más de los hashes, aunque usa Python para los ejemplos:

https://www.interviewcake.com/concept/python/hash-map

```
//Declaración  1 Forma 
temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}


// Declaración forma 2 
m:= make(map[string]int)
m["Jose"] = 25
m["Maria"] = 20

fmt.Println(m)

for i, v := range m{
	fmt.Println(i,v)
}

//Encotrar vaor 

value, ok := m[Josep]
fmt.Println(value, ok)

```
## Clase 21: Structs: La forma de hacer clases en Go

> Es una manera de generar `Clases` solo que un tantito diferente usamos la palabra reservada `Structs`

**Recordando un poco el significado de OOP:**

- Recordemos que los structs que creemos son moldes, como por ejemplo cada coche tiene una marca, tiene un color, tiene un año etc.

- Dentro del molde está el pastel y cada pastel va a ser diferente, o de manera más técnica, cada instance será diferente. Al final la clase/struct es la misma, tendremos un molde de un coche y cambiará dependiendo de las variables que mandemos.

**OOP es state y behaviour:**
- state: Son los datos
- behaviour: Son los métodos que podemos poner dentro de una clase

**Caracteristicas**
- Muy Parecido a C 
- Muy parecido a JavaScript 

**Ejemplo**
```
package main

import "fmt"

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

func main() {
	//Forma 1 de instanciar 
	myCar := car{brand: "Toyota", year: 2018, seating: 10, color: "Rojo", owner: "Eliaz Bobadilla"}
	fmt.Println("Los Datos de mi auto son:", myCar)

	//Forma Dos de instanciar 
	var otherCar car
	otherCar.brand = "Ferrari";
	otherCar.year  = 2023;

	fmt.Println("Los Datos de mi otro auto son:", otherCar)
}
```
## Clase 22: Modificadores de acceso en funciones y Structs


**Características**
- Podemos definir la variables como privada o publicas si colocamos la letra del atributo en Mayusculas ó minusculas, 
- Mayusculas -> Public
- Minusculas -> Privada
- Cuando definimos estructura `struct` Go nos indica que a fuerza debemos colocar un comentario esta sintaxis es ` // CarPublic `
- Podemos crear un mod con este comando `go mod init nombreModulo`

**Nota**
- No pude resolver esto ya que me muestra error `$ go run main.go
main.go:10:5: package basico/src/mypackage is not in GOROOT (C:\Program Files\Go\src\basico\src\mypackage)`
- No reonoce el path del paquete

## Clase 23: Structs y Punteros

**Características**
- & -> Nos permite acceder a la dirección de memoria
- * -> Nos permite acceder al valor 
- Enlace ->  https://www.digitalocean.com/community/conceptual-articles/understanding-pointers-in-go-es
```
//ejemplo 

a := 50
b := &a

fmt.Printiln(a) //-> Imprime a => 50
fmt.Printiln(b) //-> Imprime la direccion de memoria => 0xc0000a6058 
fmt.Printiln(*b) //-> Imprime el valor de la direccion de  memoria que almacena  => 50

```


## Clase 25: Interfaces y listas de interfaces
> Es un metodo que puedas compartir diferentes otros metodos, el punto es crear un solo punto de entrada. 
**Características**

- Usamos la palabra reservada `interface` este puede encerrar un metodo ya creado
- Recuerda solo toma en su estrctura una función ya creada 

```
type figure2D interface {
	getArea() float64
}


func (s square) getArea() float64 {
	return s.base * s.base
}

```

## Clase 27: ¿Qué es la concurrencia?

`
Concurrencia te permite estar pendiente de varios procesos, comienzas uno, empiezas otro, ves si el anterior ya terminó, luego crear otro así

El paralelismo, es usar varios hilos del procesador para hacer varios procesos a la vez, pero siempre estas esperando que la tarea termine.
`

**Notas**
- La concurrencia está alineando con múltiples cosas al mismo tiempo mientras que el paralelismo está haciendo múltiples cosas al mismo tiempo.
- En la concurrencia en Go podemos invocarla usando la palabra reservada en go `go say("world")`

![Ejemplo Concurrencia](../info/goland_1.png)


## Clase 28-29: Channels: La forma de organizar las goroutines

> Usamos Channels para poder manejar los gorutines
>

**Características**
- Te permite compartir datos con los gorutines 
- Goroutines sería útil a la hora de procesar documentos o imágenes.
- Solo maneja un tipo de datos 
- se usa la palabra reservada make `c := make(chan int64, 1)` 
- tenemos que usar un simbolo para ingresar un dato a ese canal `c <- int64(i)`
- tenemos que usar un simbolo para salir un dato a ese canal `<-c`
- Tenemos que usar el close en un canal luego de usarlo  si no se va usar de nuevo. 
- `Range ` es ideal cuando deseas iterar en cada uno de los elementos de un channels 
- `Select` -> Permite indicar cual goruntime va ejecutarse

```
//Salida de channels 
func printTextOurCh(c <-chan string) {
	fmt.Println(<-c)
}
```

## Clase 30: Go get: El manejador de paquetes

> Go get es el gestor de paquete para go,  

**Características**
- Tiene su pagina web 
- Pero tiene su propio motor 
- install `go get -v -u golang.org/x/tour` ó `go install -v golang.org/x/website/tour@latest`
  

**Nota**
- Como apunte adicional, he leído entre la comunidad que GO no fue pensado para manejar muchos paquetes para evitar lo que pasa con otros lenguajes como JavaScript, ya que el mayor fuerte de GO es más evidente cuando se usa el lenguaje en su estado más puro.
- 

```
Yo tengo Ubuntu 20.04.3 LTS y me sucedio esto😋:
Cuando intentaba introducir este comando en la consola:

go install -v golang.org/x/website/tour@latest
me salía este error😣:

go: modules disabled by GO111MODULE=off; see 'go help modules'

lo que hice fue introducir este comando:

export GO111MODULE=on

y luego ya pude descargarlo con el mismo comando😎:

go install -v golang.org/x/website/tour@latest
```

## Clase 31 -32 : Go modules: Ir más allá del GoPath con Echo  --- Modificando módulos con Go

> Se dio una solución en el 2018 el go modul, realiza un entorno mas facil para instalar librerias de terceros 

**Características**
- Echo es una framework para desarrollar web en GO -> https://echo.labstack.com/guide/
- instalar  -> `go get -v -u github.com/labstack/echo`
-  `go mod verify` -> Modo de verificar 
**Notas**
- Enlace -> https://github.com/golang/go/wiki/Modules


```
Recomendación para el correcto funcionamiento

Primero inicializar le modulo con go mod init myapp
Instalar el módulo con go get -u -v github.com/labstack/echo/v4
Para habilitar el módulo, debe ejecutar este comando export GO111MODULE=on
Tuve que hacerlo de esta forma, para que funcionara con WSL
```

```
Si quiren saber más sobre como mnejar los modulos en go. Les dejo este link de la documentación oficial.
https://blog.golang.org/using-go-modules

Dejó también un pequeño resumen de los comandos que se explican.

go mod init
Crea un nuevo módulo, inicializando el archivo go.mod que lo describe.
go build, go test, and other package-building commands
Agrega nuevas dependencias a go.mod según sea necesario.
go list -m all
Imprime las dependencias del módulo actual.
go get
Cambia la versión requerida de una dependencia (o agrega una nueva dependencia).
go mod tidy
Elimina las dependencias no utilizadas.
```

```
En esta lectura te llevarás algunos tópicos que siempre vale la pena tener a la mano al momento de programar en Go, además te dejo algunos tips extras.

Hola mundo
package main

import "fmt"

func main(){
    fmt.Println("Hola mundo")
}
¿Hacer una impresión en consola rápida?
package main

func main(){
    print("Hola")
}
Importar una librería sin usarla
package main

/*
    Hazlo solo y únicamente cuando la librería externa
    que estés usando lo pida explícitamente
*/
import ( 
    "fmt"
    _ "math"
)

func main(){
    fmt.Println("Hola mundo")
}
Agregar un alias a un import (no suele usarse, pero es bueno saberlo)
package main

import (
	"fmt"
	mth "math"
)

func main() {
	fmt.Println(mth.Pi)
}
Diferentes formas de declarar variables
v := 12
var v int = 12
var v int
Zero values de primitivos
var a int // 0
var b float64 // 0
var c string // ""
var d bool // false
Incremental y decremental
x++ // Suma 1 a x
x-- // Resta 1 a x
Imprimir tipo de variables (hay otras formas, pero esta es la más fácil)
a := 2
fmt.Printf("%T", a)
Función para tomar los errores (ahorra mucho código)
func isError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Ejemplo de uso
func main() {
	_, err := strconv.Atoi("53a")
	isError(err)
}
Arrays vs Slices
// Array
var myList [2]int

// Slice
var myList2 []int
Slice de interfaces (Úsalo con sabiduría)
// Permite guardar diferentes tipos de datos en un mismo slice
myList := []interface{}{"Hola", 12, 4.90}

// Iterar sobre los distintos tipos de datos de ese slice
for _, v := range myList {
    switch v.(type) {
    case int:
        fmt.Println("Es int")
    case string:
        fmt.Println("Es string")
    case float64:
        fmt.Println("Es float64")
    }
}
Asegurarnos si un key existe en el map
m := make(map[string]int)

m["hola"] = 1

// Nota, usalmente se usa "ok" para recibir la segunda variable
value, ok := m["hello"]

/*
Si existe, ok será "true"
Si no existe, ok será "false"

En este caso, ok es "false" porque no existe.
*/
Punteros
a := 10 // Variable int
b := &a // "b" es el puntero de "a"
c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"
Comandos de Go modules
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod
Nota personal
Aún tienes un largo camino por recorrer. Pero lo que más quiero que te lleves de este curso son tres cosas: Practica, estudia y participa en la comunidad de Go.

```

```
En esta lectura te llevarás algunos tópicos que siempre vale la pena tener a la mano al momento de programar en Go, además te dejo algunos tips extras.

Hola mundo
package main

import "fmt"

func main(){
    fmt.Println("Hola mundo")
}
¿Hacer una impresión en consola rápida?
package main

func main(){
    print("Hola")
}
Importar una librería sin usarla
package main

/*
    Hazlo solo y únicamente cuando la librería externa
    que estés usando lo pida explícitamente
*/
import ( 
    "fmt"
    _ "math"
)

func main(){
    fmt.Println("Hola mundo")
}
Agregar un alias a un import (no suele usarse, pero es bueno saberlo)
package main

import (
	"fmt"
	mth "math"
)

func main() {
	fmt.Println(mth.Pi)
}
Diferentes formas de declarar variables
v := 12
var v int = 12
var v int
Zero values de primitivos
var a int // 0
var b float64 // 0
var c string // ""
var d bool // false
Incremental y decremental
x++ // Suma 1 a x
x-- // Resta 1 a x
Imprimir tipo de variables (hay otras formas, pero esta es la más fácil)
a := 2
fmt.Printf("%T", a)
Función para tomar los errores (ahorra mucho código)
func isError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Ejemplo de uso
func main() {
	_, err := strconv.Atoi("53a")
	isError(err)
}
Arrays vs Slices
// Array
var myList [2]int

// Slice
var myList2 []int
Slice de interfaces (Úsalo con sabiduría)
// Permite guardar diferentes tipos de datos en un mismo slice
myList := []interface{}{"Hola", 12, 4.90}

// Iterar sobre los distintos tipos de datos de ese slice
for _, v := range myList {
    switch v.(type) {
    case int:
        fmt.Println("Es int")
    case string:
        fmt.Println("Es string")
    case float64:
        fmt.Println("Es float64")
    }
}
Asegurarnos si un key existe en el map
m := make(map[string]int)

m["hola"] = 1

// Nota, usalmente se usa "ok" para recibir la segunda variable
value, ok := m["hello"]

/*
Si existe, ok será "true"
Si no existe, ok será "false"

En este caso, ok es "false" porque no existe.
*/
Punteros
a := 10 // Variable int
b := &a // "b" es el puntero de "a"
c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"
Comandos de Go modules
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod
Nota personal
Aún tienes un largo camino por recorrer. Pero lo que más quiero que te lleves de este curso son tres cosas: Practica, estudia y participa en la comunidad de Go.
```

