package balance

import (
	"math/rand"
)

type randomAmountManager struct {
	minRange int
	maxRange int
}

func (ram randomAmountManager) generateNumber(id int) (int, error) {
	// we presume the min/max ranges are correct (max>min) as they are going to be checked in higher level modules
	return rand.Intn(ram.maxRange) + ram.minRange, nil
}

func createRandomAmountManager(minRange int, maxRange int) randomAmountManager {
	return randomAmountManager{
		minRange: minRange,
		maxRange: maxRange,
	}
}
