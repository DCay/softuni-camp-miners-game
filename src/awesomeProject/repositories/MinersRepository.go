package repositories

import (
	"awesomeProject/miners"
	"awesomeProject/minerals"
)

type MinersRepository struct {
	minersCollection []miners.BasicMiner
}

func MinersRepositoryConstruct() MinersRepository {
	return MinersRepository{
		make([]miners.BasicMiner, 0),
	}
}

func (minerRepository *MinersRepository) AddMiner() {
	minerRepository.minersCollection =
		append(minerRepository.minersCollection, miners.MinerConstruct())
}

func (minerRepository *MinersRepository) Harvest() []minerals.Mineral {
	var minedMinerals []minerals.Mineral

	for _, miner := range minerRepository.minersCollection {
		minedMinerals = append(minedMinerals, miner.Mine())
	}

	return minedMinerals
}

func (minerRepository *MinersRepository) Count() int {
	return len(minerRepository.minersCollection)
}
