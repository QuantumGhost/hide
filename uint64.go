package hide

import (
	"encoding/json"
	"math/big"
)

// Uint64 is an alias of uint64 with obfuscating/deobfuscating json marshaller
type Uint64 uint64

// MarshalJSON satisfies json.Marshaller and transparently obfuscates the value
// using Default prime
func (i *Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(Uint64Obfuscate(uint64(*i), nil))
}

// UnmarshalJSON satisfies json.Marshaller and transparently deobfuscates the
// value using coprime of Default prime
func (i *Uint64) UnmarshalJSON(data []byte) error {
	var obf uint64
	if err := json.Unmarshal(data, &obf); err != nil {
		*i = Uint64(obf)
		return err
	}
	*i = Uint64(Uint64Deobfuscate(obf, nil))
	return nil
}

// Uint64Obfuscate obfuscates uint64 provided as the 1st parameter using prime
// provided as the second one. If the provided prime is nil it will fall back
// to Default prime
func Uint64Obfuscate(val uint64, prime *big.Int) uint64 {
	if prime == nil {
		prime = Default.uint64prime
	}
	bg := new(big.Int).SetUint64(val)
	modularMultiplicativeInverse(bg, prime, uint64Max)

	return bg.Uint64()
}

// Uint64Deobfuscate deobfuscates uint64 provided as the 1st parameter using
// coprime provided as the second one. If the provided coprime is nil it will
// fall back to Default coprime
func Uint64Deobfuscate(val uint64, coprime *big.Int) uint64 {
	if coprime == nil {
		coprime = Default.uint64coprime
	}
	bg := new(big.Int).SetUint64(val)
	modularMultiplicativeInverse(bg, coprime, uint64Max)

	return bg.Uint64()
}
