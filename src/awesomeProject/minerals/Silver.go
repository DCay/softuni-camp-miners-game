package minerals

const SilverMineralName = "Silver"
const SilverMineralBasePrice = 27.50

type Silver struct {
	BaseMineral
}

func SilverConstruct() Mineral {
	return Silver {
		BaseMineral {
			SilverMineralName,
			SilverMineralBasePrice,
		},
	}
}
