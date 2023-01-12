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
- Paso 3: luego de escribir el código base o inicial debemos ir a consola y ejecutar el siguiente comando `go build ./src/main.go` este compila y genera el exe 
- Paso 4:  luego de validar que se creo el main.exe podemos ejecutarlo con el siguientes comando `./main`
- Paso 5: vemos en consola en este caso el hola mundo. 

**Otra Forma**
- Podemos ejecutar el código sin compilar pero no es recomendable 
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

**Enteras**
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