package dao

import "fmt"

type OrderDAO struct {
	db []string
}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{db: []string{}}
}

func (dao *OrderDAO) Save(order string) {
	dao.db = append(dao.db, order)
}

func (dao *OrderDAO) List() {
	fmt.Println("\n📄 Historial de pedidos:")
	for _, o := range dao.db {
		fmt.Println("- " + o)
	}
}
