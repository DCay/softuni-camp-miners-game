package minerals

const IronMineralName = "Iron"
const IronMineralBasePrice = 13.75

type Iron struct {
	BaseMineral
}

func IronConstruct() Mineral {
	return Iron {
		BaseMineral {
			IronMineralName,
			IronMineralBasePrice,
		},
	}
}
