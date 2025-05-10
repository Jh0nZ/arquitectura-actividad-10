package kitchen

import (
	"fmt"
	"sync"
)

type Kitchen struct {
	orders []string
}

var kitchenInstance *Kitchen
var once sync.Once

func GetKitchen() *Kitchen {
	once.Do(func() {
		kitchenInstance = &Kitchen{orders: []string{}}
	})
	return kitchenInstance
}

func (k *Kitchen) AddOrder(order string) {
	k.orders = append(k.orders, order)
	fmt.Println("📦 Pedido recibido en cocina:", order)
}
