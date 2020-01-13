package birthday

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"hash"
	"math"
)

// FindCollision truncates `hasher`'s output to hashLenBits and hashes
// random inputs until a collision is found
func FindCollision(hashLenBits int, hasher hash.Hash, hashName string) {
	if hashLenBits > hasher.Size()*8 {
		panic("desired length is longer than hash output size")
	}

	candidateLen := hashLenBits / 2 // giving myself plenty of room in the input space
	seen := make(map[string][]byte) // hash -> input
	hashesCreated := 0

	H := func(input []byte) string {
		hasher.Reset()
		hashesCreated++
		hasher.Write(input)
		return fmt.Sprintf("%0x", hasher.Sum(nil)[:hashLenBits/8])
	}

	random := func() []byte {
		out := make([]byte, candidateLen)
		rand.Read(out)
		return out
	}

	full := func(input []byte) string {
		hasher.Reset()
		hasher.Write(input)
		s := fmt.Sprintf("%0x", hasher.Sum(nil))
		return "[" + s[:(hashLenBits/8)*2] + "]" + s[(hashLenBits/8)*2:]
	}

	fmt.Printf("Searching for collisions in %d-bit truncated %s output\n", hashLenBits, hashName)
	fmt.Printf("Expect to find a collision in <= 2^%d/2 ~= %.f hashes\n",
		hashLenBits, math.Pow(2, float64(hashLenBits/2)),
	)
	for {
		current := random()
		h := H(current)
		if partner, found := seen[h]; found && !bytes.Equal(partner, current) {
			fmt.Printf(
				"Collision found after %d hashes generated.\n\t%0x: %s\n\t%0x: %s\n",
				hashesCreated, current, full(current), partner, full(partner),
			)
			break
		} else {
			seen[h] = current
		}
	}
}
