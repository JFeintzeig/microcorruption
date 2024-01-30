package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	h := sha256.New()

	hashes := make(map[string]uint8)

	// loop over all characters, maybe all bigrams?
	for i := uint16(0); i < 256; i++ {
		h.Write([]byte{uint8(i)})
		hash := h.Sum(nil)
		hashString := hex.EncodeToString(hash)
		hashes[hashString[:8]] = uint8(i)
		h.Reset()
	}

	file, err := os.ReadFile("printed_hashes.txt")
	if err != nil {
		panic("problem opening file")
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if val, ok := hashes[strings.ToLower(line)]; ok {
			fmt.Printf("%02x", val)
		} else {
			fmt.Printf("\nProblem: no hash found for %s\n", line)
		}
	}
	fmt.Printf("\n")
}
