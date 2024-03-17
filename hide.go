package hide

import (
	"math"
	"math/big"
)

type sentinel string

func (s sentinel) Error() string {
	return string(s)
}

var (
	bigOne = big.NewInt(1)

	// maximum value for each type

	ErrNotAPrime    sentinel = "it is not a prime number"
	ErrInvalidPrime sentinel = "invalid prime (cannot use 2 as prime)"
)

type Unsigned interface {
	uint32 | uint64 | uint
}

type Hide[T Unsigned] struct {
	prime    *big.Int
	inverse  *big.Int
	xor      *big.Int
	maxValue *big.Int
}

func New[T Unsigned](prime T, xor T) (*Hide[T], error) {
	bigPrime := new(big.Int).SetUint64(uint64(prime))
	if !bigPrime.ProbablyPrime(100) {
		return nil, ErrNotAPrime
	}
	if bigPrime.Cmp(big.NewInt(2)) == 0 {
		return nil, ErrInvalidPrime
	}
	maxValue := maxOf[T]()
	hider := &Hide[T]{
		prime:    bigPrime,
		inverse:  getInverse(bigPrime, maxValue),
		maxValue: maxValue,
	}

	if xor != 0 {
		hider.xor = new(big.Int).SetUint64(uint64(xor))
	}
	return hider, nil
}

func (h *Hide[T]) Obfuscate(value T) T {
	bg := new(big.Int).SetUint64(uint64(value))
	modularMultiplicativeInverse(bg, h.prime, h.maxValue)

	if xor := h.xor; xor != nil {
		bg.Xor(bg, xor)
	}

	return T(bg.Uint64())
}

func (h *Hide[T]) Deobfuscate(value T) T {
	bg := new(big.Int).SetUint64(uint64(value))
	if xor := h.xor; xor != nil {
		bg.Xor(bg, xor)
	}

	modularMultiplicativeInverse(bg, h.inverse, h.maxValue)

	return T(bg.Uint64())
}

func getInverse(prime *big.Int, maxValue *big.Int) *big.Int {
	var inverse big.Int
	inverse.Set(prime)

	var plusOne big.Int
	return inverse.ModInverse(&inverse, plusOne.Add(maxValue, bigOne))
}

func maxOf[T Unsigned]() *big.Int {
	switch any(T(0)).(type) {
	case uint32:
		return big.NewInt(math.MaxUint32)
	case uint64:
		return new(big.Int).SetUint64(math.MaxUint64)
	case uint:
		return new(big.Int).SetUint64(math.MaxUint)
	default:
		// make compiler happy.
		panic("unreachable")
	}
}

func modularMultiplicativeInverse(val, prime, max *big.Int) {
	val.Mul(val, prime)
	val.And(val, max)
}
