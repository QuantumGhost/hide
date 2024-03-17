package hide_test

import (
	"fmt"
	"github.com/QuantumGhost/hide"
	"math/rand/v2"
)

func Example() {

	hider, err := hide.New[uint32](65537, 2813448142)
	if err != nil {
		panic(err)
	}

	fmt.Println("Random IDs")
	for i := uint32(0); i < 10; i++ {
		v := rand.N[uint32](1000000)

		o := hider.Obfuscate(v)

		fmt.Printf("%8d -> %10d\n", v, o)
	}

	fmt.Println("\nConsecutive IDs")
	start := rand.N[uint32](1000000)
	for i := start; i < start+10; i++ {
		o := hider.Obfuscate(i)

		fmt.Printf("%8d -> %10d\n", i, o)
	}
}
