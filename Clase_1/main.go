package main

import "fmt"

/*
const miConstante = "Valor de mi constante ñandú"

//variables y constantes
func main() {
	//declaración por inferencia
	var nombre string = "Cesar"
	nombre1 := "Cesar"
	fmt.Println(nombre)
	fmt.Print(nombre1)
	//declaracion rapida o corta
	nombre2 := "Juan"
	fmt.Println(nombre2)
	fmt.Printf("El valor de miConstante es %s \n", miConstante)
}
*/
/*
//tipos de datos
func main() {
	var string1 string = "texto"
	fmt.Println(string1)
	textoGrande := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	fmt.Println(textoGrande)
	var estado bool = true
	fmt.Println(estado)
	var entero int8 = 123
	fmt.Println(entero)
	var entero1 int64 = 78910
	fmt.Println(entero1)
}
*/
/*
//reflect y typeOf
func main() {
	var string1 string = "texto"
	fmt.Println(reflect.TypeOf(string1))
	fmt.Println((reflect.TypeOf(string1)))
}
*/
/*
// punteros
func main() {
	color := "rojo"
	hijo := "Eustaquio"
	fmt.Println(color)
	fmt.Println(hijo)
	fmt.Println(color, &color)
	var estado bool = true
	fmt.Println(&estado)
}
*/
/*
//condicionales
func main() {
	edad := 30
	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
	fmt.Println("-------------------------------")
	color := "blanco"
	if color == "rojo" {
		fmt.Println("Es rojo como la sangre de los araucamos")
	} else if color == "azul" {
		fmt.Println("Es azul como el cielo")
	} else {
		fmt.Println("Es blanco como la nieve")
	}
	fmt.Println("-------------------------")
	if color == "azul" && edad == 11 {
		fmt.Println("color es azul y edad 11")
	}
	fmt.Println("-------------------------")
	if variable := 2; variable == 2 {
		fmt.Println("Variable es igual a 2")
	}
	//switch case
	switch color {
	case "rojo":
		fmt.Println("Es de color rojo")

	case "azul":
		fmt.Println("El color es azul")
	}
}
*/
/*
import "fmt"

func main() {
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}
	fmt.Println("------------------------")
	for i2 := 1; i2 < 11; i2 += 3 {
		fmt.Println(i2)
		if i2 == 6 {
			continue
		}
		if i2%2 == 0 {
			fmt.Println("Es par")
		}
	}
}
*/
/*
//arreglos y slice
func main() {
	//arreglo
	var paises [4]string
	paises[0] = "Chile"
	paises[1] = "Argentina"
	paises[2] = "Peru"
	paises[3] = "Colombia"
	fmt.Println(paises)
	fmt.Println(paises[1])
	fmt.Println("El largo del arreglo es ", len(paises))
	//slice
	paises2 := make([]string, 5)
	paises2[0] = "Chile"
	//agregar un elemento al slice
	paises2 = append(paises2, "Argentina")
	fmt.Println(paises2)
	paises2 = append(paises2[:5], paises2[5+1:]...)
	fmt.Println(paises2)
}
*/
/*
// map
func main() {
	paises := make(map[string]int)
	paises["argentina"] = 4000000
	paises["españa"] = 4500000
	paises["chile"] = 18000000
	fmt.Println("\n", paises)
	fmt.Println(paises["chile"])
	paises2 := map[int]string{
		1: "Chile",
		2: "Argentina",
		3: "España",
	}
	fmt.Println(paises2)
	fmt.Println(paises2[1])
	//Veamos si existe algun valor en el map
	pais, existe := paises2[11]
	if existe {
		fmt.Println("Si existe el país", pais)
	} else {
		fmt.Println("No existe el país", pais)
	}
	delete(paises2, 1)
	fmt.Println(paises2)
}
*/
/*
//funciones
func main() {
	miFuncion()
	miFuncionConParametros(10, 2)
	fmt.Println(miFuncionConRetorno("Adrián"))
	fmt.Println(miFuncionConRetornoMultiple())
	fmt.Println(suma(1, 2))
	fmt.Println(tabla()())
}

func miFuncion() {
	fmt.Println("¡Hola Mundo!")
}

func miFuncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("La suma es :", resultado)
}

func miFuncionConRetorno(nombre string) string {
	return "tu nombre es " + nombre
}

func miFuncionConRetornoMultiple() (string, string, int) {
	return "Adrián", "Gil", 31
}

//funciones anonimas
var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

//funciones clousure
func tabla() func() int {
	numero := 1 // Initialize with a default value since 'valor' is undefined
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
*/
/*
// gorutinas
func main() {
	//ejemplo1
	fmt.Println(miFuncion("Hola"))
	time.Sleep(time.Second * 5)
	fmt.Println("Juana")
	//ejemplo2
	miCanal := make(chan string)
	go func() {
		miCanal <- miFuncion("Pedro")
	}()
	fmt.Println(<-miCanal)
	fmt.Println("Continuar con la ejecución")
}

func miFuncion(parametro string) string {
	return "hola " + parametro
}
*/
//recursividad
/*
func main() {
	miFuncion(0)
}

func miFuncion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncion(dato)
	}
}
*/
/*
//defer y panic
import "fmt"

func main() {
	miFuncion()
}

func miFuncion() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("este es un primer mensaje")
	a := 1
	if a == 1 {
		panic("Fallo porque falló")
	}
}
*/

/*
// POO
func main() {
	estructura := Persona{
		Id:     1,
		Nombre: "Adrián Gil",
		Correo: "info@tamila.cl",
		Edad:   42,
	}
	fmt.Println(estructura)
	fmt.Println(reflect.TypeOf(estructura))
	p := new(Persona)
	p.Id = 1
	p.Nombre = "Adrián Gil"
	fmt.Println(p.Nombre)
}

type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}
type Categoria struct {
    Id int
    Nombre: "Categoria 1",
    Slug: "categoria-1",
}
producto := Producto(Id:1, Nombre: "Mesa de computador", Slug: "mesa-de-computador")
fmt.Println(producto)
fmt.Println(categoria)
//estructura anidada
*/

//Interfaces
func main() {
	e := Estructura{}
	fmt.Println(e.miFuncion())
	fmt.Println(e.Calculo(1, 2))
}

type EjemploInterface interface {
	miFuncion() string
	Calculo(n1 int, n2 int) int
}

type Estructura struct {
}

func (*Estructura) miFuncion() string {
	return "texto desde mí función"
}

func (*Estructura) Calculo(n1 int, n2 int) int {
	resultado := n1 + n2
	return resultado
}
