package data

import (
	"fastfood/application"
)

type RepositorioPedidos struct {
	db []application.Pedido
}

func NuevoRepositorio() *RepositorioPedidos {
	return &RepositorioPedidos{
		db: []application.Pedido{},
	}
}

func (r *RepositorioPedidos) Guardar(p application.Pedido) {
	r.db = append(r.db, p)
}

func (r *RepositorioPedidos) Listar() []application.Pedido {
	return r.db
}
