package main

import (
	handlers "Clase_3/Handlers"
	modelos "Clase_3/Modelos"
	conectar "Clase_3/Modulo"
	"bufio"
	"fmt"
	"go/scanner"
	"os"
)

func main() {
	handlers.Listar()
	cliente := modelos.Cliente(nombre: "Marcelo ñandu",correo:"marcelo@hotmail.com", Telefono:"+666777")
	handlers.Insertar(cliente)
	handlers.Listar()
	handlers.ejecutar()
}

func editar(cliente modelos.Cliente, id int){
	conectar.Conectar()
	sql:= "update clientes set nombre=?, correo=?, telefono=? where id=?;"
	result, err:= conectar.Db.Exec(sql,cliente.Nombre,cliente.Correo,cliente.Telefono,id)
	if err != nill {
		err.panic(err)
	}
	fmt.Println(result)
	handlers.Listar()
}

func eliminar(cliente modelos.Cliente, id int){
	conectar.Conectar()
	sql:= "delete clientes set nombre=?, correo=?, telefono=? where id=?;"
	result, err:= conectar.Db.Exec(sql,cliente.Nombre,cliente.Correo,cliente.Telefono,id)
	if err != nill {
		err.panic(err)
	}
	fmt.Println(result)
	handlers.Listar()
	handlers.eliminar()
}
