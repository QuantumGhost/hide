package hide_test

import (
	"fmt"
	"github.com/QuantumGhost/hide"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"pgregory.net/rapid"
	"testing"
)

func integerDrawer[V hide.Unsigned](t *rapid.T) V {
	switch any(V(0)).(type) {
	case uint32:
		return V(rapid.Uint32().Draw(t, "uint32"))
	case uint64:
		return V(rapid.Uint64().Draw(t, "uint64"))
	case uint:
		return V(rapid.Uint().Draw(t, "uint"))
	default:
		panic("unreachable")
	}
}

func runRapidTestWith[V hide.Unsigned](t rapid.TB) {
	rapid.Check(t, func(t *rapid.T) {
		prime := rapid.Custom[V](integerDrawer[V]).Filter(func(v V) bool {
			if v < 3 {
				return false
			}
			return new(big.Int).SetUint64(uint64(v)).ProbablyPrime(100)
		}).Draw(t, "prime")
		xor := rapid.Custom[V](integerDrawer[V]).Draw(t, "xor")
		hider, err := hide.New(prime, xor)
		assert.NoError(t, err)
		assert.NotNil(t, hider)
		for i := 0; i < 1000; i++ {
			value := rapid.Custom[V](integerDrawer[V]).Draw(t, fmt.Sprintf("value[%d]", i))
			require.Equal(t, value, hider.Deobfuscate(hider.Obfuscate(value)))
		}
	})
}

func TestWithRandomValue(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		runRapidTestWith[uint32](t)
	})
	t.Run("uint64", func(t *testing.T) {
		runRapidTestWith[uint64](t)
	})
	t.Run("uint", func(t *testing.T) {
		runRapidTestWith[uint64](t)
	})
}

func FuzzHideUint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, prime uint32, xor uint32, value uint32) {
		if prime < 3 {
			t.Skip()
		}
		if value < 0 {
			t.Skip()
		}
		if !big.NewInt(int64(prime)).ProbablyPrime(100) {
			t.Skip()
		}

		hider, err := hide.New(prime, xor)
		assert.NoError(t, err)
		assert.NotNil(t, hider)
		assert.Equal(t, value, hider.Deobfuscate(hider.Obfuscate(value)))
	})
}

func FuzzHideUint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, prime uint64, xor uint64, value uint64) {
		if prime < 3 {
			t.Skip()
		}
		if value < 0 {
			t.Skip()
		}
		if !new(big.Int).SetUint64(prime).ProbablyPrime(100) {
			t.Skip()
		}

		hider, err := hide.New(prime, xor)
		assert.NoError(t, err)
		assert.NotNil(t, hider)
		assert.Equal(t, value, hider.Deobfuscate(hider.Obfuscate(value)))
	})
}

func FuzzHideUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, prime uint, xor uint, value uint) {
		if prime < 3 {
			t.Skip()
		}
		if value < 0 {
			t.Skip()
		}
		if !new(big.Int).SetUint64(uint64(prime)).ProbablyPrime(100) {
			t.Skip()
		}

		hider, err := hide.New(prime, xor)
		assert.NoError(t, err)
		assert.NotNil(t, hider)
		assert.Equal(t, value, hider.Deobfuscate(hider.Obfuscate(value)))
	})
}
