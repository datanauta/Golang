package modelos

type Cliente struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
}

type Clientes []Cliente
