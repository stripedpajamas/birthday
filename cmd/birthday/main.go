package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"os"

	"github.com/stripedpajamas/birthday/internal/birthday"
)

// the expected number of N-bit hashes that can be generated
// before getting a collision is not 2^N, but rather only 2​^N⁄2
func main() {
	bits := flag.Int("bits", 40, "bit length")
	alg := flag.String("alg", "md5", "hashing algorithm to use (md5, sha1, sha256, sha512)")
	flag.Parse()

	if *bits == 0 || *alg == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var hasher hash.Hash
	switch *alg {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		fmt.Println("Supported hashing algorithms: md5, sha1, sha256, sha512.")
		os.Exit(1)
	}

	birthday.FindCollision(*bits, hasher, *alg)
}
