package hide

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

func TestInternal(t *testing.T) {
	assert.Equal(t, big.NewInt(1), bigOne)
	assert.Equal(t, big.NewInt(math.MaxUint32), maxOf[uint32]())
	assert.Equal(t, new(big.Int).SetUint64(math.MaxUint64), maxOf[uint64]())
	assert.Equal(t, new(big.Int).SetUint64(math.MaxUint), maxOf[uint]())
}
