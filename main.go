package main

import (
	"fastfood/dao"
	"fastfood/delivery"
	"fastfood/food"
	"fastfood/kitchen"
)

func main() {
	dao := dao.NewOrderDAO()
	k := kitchen.GetKitchen()

	// Productos a pedir
	orders := []string{"burger", "pizza", "hotdog", "nuggets", "tacos"}
	deliveries := []delivery.DeliveryStrategy{
		delivery.HomeDelivery{},
		delivery.InStorePickup{},
	}

	for i, item := range orders {
		f := food.FoodFactory(item)
		if f == nil {
			continue
		}
		result := f.Prepare()

		k.AddOrder(result)
		dao.Save(result)

		// Alternar entre entregas
		deliveries[i%2].Deliver(result)
	}

	dao.List()
}
