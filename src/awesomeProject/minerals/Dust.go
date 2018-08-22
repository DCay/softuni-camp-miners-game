package minerals

const DustMineralName = "Dust"
const DustMineralBasePrice = 0

type Dust struct {
	BaseMineral
}

func DustConstruct() Mineral {
	return Dust {
		BaseMineral {
			GoldMineralName,
			GoldMineralBasePrice,
		},
	}
}