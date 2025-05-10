package food

// Product: definición de las interfaces para la familia de productos genéricos.
type Food interface {
	Prepare() string
}

// ConcreteProduct: implementación de los diferentes productos.
type Burger struct{}
func (b Burger) Prepare() string { return "🍔 Hamburguesa lista" }

type Pizza struct{}
func (p Pizza) Prepare() string { return "🍕 Pizza lista" }

type HotDog struct{}
func (h HotDog) Prepare() string { return "🌭 Hot Dog listo" }

type Nuggets struct{}
func (n Nuggets) Prepare() string { return "🍗 Nuggets listos" }

type Tacos struct{}
func (t Tacos) Prepare() string { return "🌮 Tacos listos" }

// ConcreteProduct: implementación de los diferentes productos.
func FoodFactory(foodType string) Food {
	switch foodType {
	case "burger":
		return Burger{}
	case "pizza":
		return Pizza{}
	case "hotdog":
		return HotDog{}
	case "nuggets":
		return Nuggets{}
	case "tacos":
		return Tacos{}
	default:
		return nil
	}
}
