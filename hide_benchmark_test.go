package hide_test

import (
	"github.com/QuantumGhost/hide"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

func BenchmarkUint32Obfuscate(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, big.NewInt(math.MaxUint32))
	hider, err := hide.New[uint32](uint32Prime, uint32(bigXor.Int64()))
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint32(i))
	}
}

func BenchmarkUint32ObfuscateNoXor(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, big.NewInt(math.MaxUint32))
	hider, err := hide.New[uint32](uint32Prime, 0)
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint32(i))
	}
}

func BenchmarkUint64Obfuscate(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	hider, err := hide.New[uint64](uint64Prime, bigXor.Uint64())
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint64(i))
	}
}

func BenchmarkUint64ObfuscateNoXor(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	hider, err := hide.New[uint64](uint64Prime, 0)
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint64(i))
	}
}

func BenchmarkUintObfuscate(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	hider, err := hide.New[uint](uint64Prime, uint(bigXor.Uint64()))
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint(i))
	}
}

func BenchmarkUintObfuscateNoXor(b *testing.B) {
	bigXor := big.NewInt(xorValue)
	bigXor.And(bigXor, new(big.Int).SetUint64(math.MaxUint64))
	hider, err := hide.New[uint](uint64Prime, 0)
	assert.NoError(b, err)
	for i := 1; i <= b.N; i++ {
		hider.Obfuscate(uint(i))
	}
}
