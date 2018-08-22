package miners

import (
	"math/rand"
	"awesomeProject/minerals"
)

//COAL - 45%
//COPPER - 30%
//IRON - 20%
//SILVER - 4%
//GOLD - 1%

const BasicMinerBaseRpm = 2
const BasicMinerBaseOutput = 1
const BasicMinerBaseLevel = 1

type BasicMiner struct {
	rpm    int
	output int
	level  int
}

func MinerConstruct() BasicMiner {
	return BasicMiner{
		BasicMinerBaseRpm,
		BasicMinerBaseOutput,
		BasicMinerBaseLevel,
	}
}

func (miner BasicMiner) GetRpm() int {
	return miner.rpm + (miner.level - 1)
}

func (miner BasicMiner) GetOutput() int {
	return miner.output + (miner.level - 1)
}

func (miner BasicMiner) GetLevel() int {
	return miner.level
}

func (miner BasicMiner) Mine() minerals.Mineral {
	chance := rand.Intn(100);

	if chance > 55 && chance <= 100 {
		return minerals.CoalConstruct()
	} else if chance > 25 && chance <= 55 {
		return minerals.CopperConstruct()
	} else if chance > 5 && chance <= 25 {
		return minerals.IronConstruct()
	} else if chance > 1 && chance <= 5 {
		return minerals.SilverConstruct()
	} else if chance == 1 {
		return minerals.GoldConstruct()
	} else {
		return minerals.DustConstruct()
	}
}
