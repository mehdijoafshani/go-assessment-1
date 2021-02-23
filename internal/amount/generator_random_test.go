package amount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomAmountManagerGenerateNumber(t *testing.T) {
	min := 1
	max := 2
	randomAmountManager := createRandomAmountManager(min, max)

	for i := 0; i < 100; i++ {
		generatedNumber, err := randomAmountManager.generateNumber(0)
		assert.Nil(t, err, "the method should return no error")
		assert.GreaterOrEqual(t, generatedNumber, min)
		assert.LessOrEqual(t, generatedNumber, max)
	}
}
