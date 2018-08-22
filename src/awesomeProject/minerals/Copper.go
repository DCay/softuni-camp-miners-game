package minerals

const CopperMineralName = "Copper"
const CopperMineralBasePrice = 7.5

type Copper struct {
	BaseMineral
}

func CopperConstruct() Mineral {
	return Copper {
		BaseMineral {
			CopperMineralName,
			CopperMineralBasePrice,
		},
	}
}