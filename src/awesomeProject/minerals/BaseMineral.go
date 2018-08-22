package minerals

type BaseMineral struct {
	name string
	price float64
}

func (mineral BaseMineral) GetName() string {
	return mineral.name
}

func (mineral BaseMineral) GetPrice() float64 {
	return mineral.price
}