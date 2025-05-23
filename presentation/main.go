package main

import (
	"fastfood/application"
	"fastfood/data"
	"fmt"
)

func main() {
	repo := data.NuevoRepositorio()
	servicio := application.NuevoServicio(repo)

	cliente := application.Cliente{Nombre: "Ion"}

	servicio.RegistrarPedido(cliente, "Pizza", application.EnTienda{})
	servicio.RegistrarPedido(cliente, "Burger", application.ADomicilio{})

	fmt.Println("\n📄 Historial de pedidos:")
	for _, p := range servicio.ListarPedidos() {
		fmt.Printf("- %s pidió %s (%s)\n", p.Cliente.Nombre, p.Producto.Nombre, p.Entrega)
	}
}
