package application

type Repositorio interface {
	Guardar(p Pedido)
	Listar() []Pedido
}
