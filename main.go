package main

import (
	"encoding/hex"
	"fmt"
  "math/rand"
)

// chernobyl hash collisions

func sxt(b byte) uint16 {
  if ((b >> 7) & 0x1) == 0x1 {
    return uint16(0xff00) | uint16(b)
  } else {
    return uint16(b)
  }
}

func hash(input []byte) uint16 {
  //fmt.Printf("\n\nstart hash\n")
  var running uint16
  for _, b := range input {
    temp := sxt(b)
    temp += running
    running = temp
    running *= 2
    running *= 2
    running *= 2
    running *= 2
    running *= 2
    running += (^temp + 1)
    //fmt.Printf("hash -- input:%x running:%x\n", b, running)
  }
  return running
}

func mask(input uint16, masknum int) uint16 {
  //fmt.Printf("mask: %x\n", (1 << masknum - 1))
  return input & ((1 << masknum) - 1)
}

func GenerateRandomBytes(length int) []byte {
	out := make([]byte, length)
	for i := range out {
		out[i] = uint8(rand.Intn(255))
	}
	return out
}

func main() {
  cases := [][]byte{[]byte("abcdef"), []byte("qwertyqwerty"), []byte("123456789101112")}
  for _, input := range cases {
    fmt.Printf("******* Input: %s *********\n", input)
    output := hash(input)
    fmt.Printf("hash: %x\n", output)
    masked3 := mask(output, 3)
    masked4 := mask(output, 4)
    fmt.Printf("masked 3, 4: %x, %x\n", masked3, masked4)
  }

  test, err := hex.DecodeString("3c538856b5018679491a712ebd734940")
  if err != nil {
    panic("rpbo")
  }

  fmt.Printf("payload: %x %x %d %d\n", test, hash(test), mask(hash(test),3), mask(hash(test),4))

  seed, err := hex.DecodeString("0105dc52b5010a50e450cd02d248d248d2")
  if err != nil {
    panic("problem decoding hex string")
  }

  found := 0
  for true {
    rBytes := GenerateRandomBytes(1)
    pl := append(seed, rBytes...)
    bin3 := mask(hash(pl), 3)
    bin4 := mask(hash(pl), 4)
    if bin3 == 1 {
      found += 1
      fmt.Printf("payload %x %d %d\n", pl, bin3, bin4)
    }
    if found == 10 {
      break
    }
  }
}
