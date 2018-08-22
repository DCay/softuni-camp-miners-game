package minerals

const CoalMineralName = "Coal"
const CoalMineralBasePrice = 5.0

type Coal struct {
	BaseMineral
}

func CoalConstruct() Mineral {
	return Coal {
		BaseMineral {
			CoalMineralName,
			CoalMineralBasePrice,
		},
	}
}