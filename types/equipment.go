package types

type Equipment struct {
	name  string
	price int
}

func NewPickaxes() *Equipment {
	return &Equipment{
		name:  "pickaxes",
		price: 3000,
	}
}

func NewVentilation() *Equipment {
	return &Equipment{
		name:  "ventilation",
		price: 15000,
	}
}

func NewCars() *Equipment {
	return &Equipment{
		name:  "cars",
		price: 50000,
	}
}

func (eqiup *Equipment) GetData() (string, int) {
	return eqiup.name, eqiup.price
}