package minerals

const GoldMineralName = "Gold"
const GoldMineralBasePrice = 75.25

type Gold struct {
	BaseMineral
}

func GoldConstruct() Mineral {
	return Gold {
		BaseMineral {
			GoldMineralName,
			GoldMineralBasePrice,
		},
	}
}