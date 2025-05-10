package delivery

import "fmt"

type DeliveryStrategy interface {
	Deliver(food string)
}

type InStorePickup struct{}
func (d InStorePickup) Deliver(food string) {
	fmt.Println("🏪 Cliente recoge el pedido:", food)
}

type HomeDelivery struct{}
func (d HomeDelivery) Deliver(food string) {
	fmt.Println("🚚 Repartidor entrega el pedido:", food)
}
