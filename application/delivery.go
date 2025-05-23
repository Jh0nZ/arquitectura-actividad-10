package application

import "fmt"

type DeliveryStrategy interface {
	Entregar(pedido Pedido)
	Nombre() string
}

type EnTienda struct{}
func (e EnTienda) Entregar(p Pedido) {
	fmt.Printf("🏪 %s recoge su %s en tienda\n", p.Cliente.Nombre, p.Producto.Nombre)
}
func (e EnTienda) Nombre() string { return "En tienda" }

type ADomicilio struct{}
func (d ADomicilio) Entregar(p Pedido) {
	fmt.Printf("🚚 Se entrega %s a %s en su domicilio\n", p.Producto.Nombre, p.Cliente.Nombre)
}
func (d ADomicilio) Nombre() string { return "A domicilio" }
