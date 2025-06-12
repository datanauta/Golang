package main

/*
// modulos
func main() {
	fmt.Println("La hora actual es : ", time.Now())
	fecha := time.Now()
	fmt.Println(fecha.Day())
	fmt.Println(fecha.Year())
	fmt.Println(fecha.Month())
	fmt.Println(fecha.Minute())
	fmt.Println(fecha.Second())
	ahora := time.Now()
	fecha1 := ahora.Add(time.Hour * 20 * 24)
	fmt.Println(fecha1)
}
*/
/*
func main() {
	cadena := "mi muñeca me hablo"
	fmt.Println(cadena)
	fmt.Println(strings.ToLower(cadena))
	pos := strings.Index(cadena, "Mi")
	if pos == -1 {
		fmt.Println("No se encontró")
	} else {
		fmt.Println("Se encontró en la posición")
	}
	repetida := strings.Repeat(cadena, 2)
	fmt.Println(repetida)
	cadenanueva := strings.Replace(cadena, "Mi", "la", -1)
	fmt.Println(cadenanueva)
	fmt.Println(string(cadenanueva[0:5]))
	fmt.Println(string(cadenanueva[0:9]))
}
*/
/*
//math rand
func main() {
	aleatorio := rand.Intn(101)
	fmt.Println(aleatorio)
	min := 1000
	max := 10000
	rand.Seed(time.Now().UnixNano())
	aleatorio2 := rand.Intn(max-min) + min
	fmt.Println(aleatorio2)
}
*/
/*
//modulo os
func main() {
	nombre := flag.String("Nombre", "", "El nombre de la persona")
	edad := flag.Int("edad", 18, "La edad de la persona")
	flag.Parse()
	fmt.Println("Tu nombre es:", *nombre)
	fmt.Println("Tu edad es:", *edad)
}
*/
//modulo log

/*
func main() {
	err := errors.New("Este es un error fatal de prueba")
	log.Fatal(err)
	log.Fatal("Error fatal")
	log.Panic("Error Panic")
	log.Panicln("Error Panic")

}
*/
