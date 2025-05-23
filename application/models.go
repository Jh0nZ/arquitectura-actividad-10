package application

type Cliente struct {
	Nombre string
}

type Producto struct {
	Nombre string
}

type Pedido struct {
	Cliente  Cliente
	Producto Producto
	Entrega  string
}
