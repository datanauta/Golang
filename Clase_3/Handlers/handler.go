package handlers

import (
	modelos "Clase_3/Modelos"
	"fmt"

	"github.com/Usuario/Clase_3/mysql_driver/conectar"
)

func Listar() {
	fmt.Println("Hola")

	// Conexión a la base de datos
	conectar.Conectar()
	defer conectar.Db.Close() // Cerramos conexión al final

	// Consulta SQL
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return
	}
	defer datos.Close()

	// Creamos slice para guardar los resultados
	var clientes []modelos.Cliente

	// Iteramos sobre los resultados
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			fmt.Println("Error al escanear fila:", err)
			continue
		}
		clientes = append(clientes, dato)
	}

	// Verificamos si hubo errores al iterar
	if err := datos.Err(); err != nil {
		fmt.Println("Error al iterar los resultados:", err)
	}

	// Mostramos los datos
	for _, c := range clientes {
		fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Teléfono: %s\n", c.Id, c.Nombre, c.Correo, c.Telefono)
	}

	fmt.Println("Lstado de clientes")
	for datos.Next() {
		var dato = modelos.Cliente{
			fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Teléfono: %s\n", c.Id, c.Nombre, c.Correo, c.Telefono)
		}
	}
}
funct ListarPorId(id int){
	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes where id =?;"
	datos,err := conectar.Db.Query(sql,id)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
}
func Insertar(cliente modelos.Cliente){
	conectar.Conectar()
	sql:="insert into clientes values(null,?,?,?)"
	result, err:= conectar.Db.Exec(sql,cliente.Nombre,cliente.Correo, cliente.Telefono)
	if err != nill{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println("Se creó el registro exitosamente")
}

