package application

type ServicioPedido struct {
	repo Repositorio
}

func NuevoServicio(repo Repositorio) *ServicioPedido {
	return &ServicioPedido{repo}
}

func (s *ServicioPedido) RegistrarPedido(cliente Cliente, nombreProducto string, entrega DeliveryStrategy) {
	producto := CrearProducto(nombreProducto)
	pedido := Pedido{
		Cliente:  cliente,
		Producto: producto,
		Entrega:  entrega.Nombre(),
	}
	s.repo.Guardar(pedido)
	entrega.Entregar(pedido)
}

func (s *ServicioPedido) ListarPedidos() []Pedido {
	return s.repo.Listar()
}
