package main

import (
  "fmt"
  "math/rand"
)

func GenerateRandomWords(length int) []uint16 {
	out := make([]uint16, length)
	for i := range out {
		out[i] = uint16(rand.Intn(0xFFFF))
	}
	return out
}

func swap(x uint16) uint16 {
  return ((x & 0xff) << 8) | ((x & 0xff00) >> 8)
}

func CalcStuff(input []uint16) {
  var r4 uint16 = 0
  var r6 uint16 = 0
  for i, val := range input {
    r4 += swap(val)
    r4 = swap(r4)
    r6 ^= swap(val)
    r6 ^= r4
    r4 ^= r6
    r6 ^= r4
    if r6 == 0x9298 && r4 == 0xFEB1 {
      fmt.Printf("%v %d", input, i)
      panic("done")
    }
  }
}

func main() {

  //CalcStuff([]uint16{0x6162, 0x6364, 0x6566})

  i := 0
  for true {
    if i % 1e7 == 0 {
      fmt.Printf("%d\n", i/1e7)
    }
    input := GenerateRandomWords(8)
    CalcStuff(input)
    i+=1
  }
}
