package hide_test

import (
	"github.com/QuantumGhost/hide"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"math/rand"
	"testing"
)

const (
	uint32Prime = 1500450271
	uint64Prime = 12764787846358441471
	xorValue    = 3469983624777167712
)

func TestUint32Reversible(t *testing.T) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, big.NewInt(math.MaxUint32))
	xorValues := []uint32{uint32(bigXor.Uint64()), 0}
	for _, xor := range xorValues {
		hider, err := hide.New[uint32](uint32Prime, xor)
		assert.NoError(t, err)
		for i := 0; i < 100000; i++ {
			v := uint32(rand.Int31() * (1 + rand.Int31n(1)))
			r := hider.Deobfuscate(hider.Obfuscate(v))
			if v != r {
				t.Logf("Expected %d, actual %d, xor=%d", v, r, xor)
				t.Fail()
			}
		}
	}
}

func TestUint64Reversible(t *testing.T) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	xorValues := []uint64{bigXor.Uint64(), 0}
	for _, xor := range xorValues {
		hider, err := hide.New[uint64](uint64Prime, xor)
		assert.NoError(t, err)
		for i := 0; i < 100000; i++ {
			v := uint64(rand.Int63() * (1 + rand.Int63n(1)))
			r := hider.Deobfuscate(hider.Obfuscate(v))
			if v != r {
				t.Logf("Expected %d, actual %d, xor=%d", v, r, xor)
				t.Fail()
			}
		}
	}
}

func TestUintReversible(t *testing.T) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	xorValues := []uint{uint(bigXor.Uint64()), 0}
	for _, xor := range xorValues {
		hider, err := hide.New[uint](uint64Prime, xor)
		assert.NoError(t, err)
		for i := 0; i < 100000; i++ {
			v := uint(rand.Int63() * (1 + rand.Int63n(1)))
			r := hider.Deobfuscate(hider.Obfuscate(v))
			if v != r {
				t.Logf("Expected %d, actual %d, xor=%d", v, r, xor)
				t.Fail()
			}
		}
	}
}

func TestNewWithNonPrime(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		hider, err := hide.New[uint32](1<<31, 0)
		assert.ErrorIs(t, err, hide.ErrNotAPrime)
		assert.Nil(t, hider)
	})
	t.Run("uint64", func(t *testing.T) {
		hider, err := hide.New[uint64](1<<63, 0)
		assert.ErrorIs(t, err, hide.ErrNotAPrime)
		assert.Nil(t, hider)
	})
	t.Run("uint", func(t *testing.T) {
		hider, err := hide.New[uint](1<<63, 0)
		assert.ErrorIs(t, err, hide.ErrNotAPrime)
		assert.Nil(t, hider)
	})
}

func TestNewWithPrime2(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		hider, err := hide.New[uint32](2, 0)
		assert.ErrorIs(t, err, hide.ErrInvalidPrime)
		assert.Nil(t, hider)
	})
	t.Run("uint64", func(t *testing.T) {
		hider, err := hide.New[uint64](2, 0)
		assert.ErrorIs(t, err, hide.ErrInvalidPrime)
		assert.Nil(t, hider)
	})
	t.Run("uint", func(t *testing.T) {
		hider, err := hide.New[uint](2, 0)
		assert.ErrorIs(t, err, hide.ErrInvalidPrime)
		assert.Nil(t, hider)
	})
}

func TestSentinelError(t *testing.T) {
	assert.Equal(t, "it is not a prime number", hide.ErrNotAPrime.Error())
}
