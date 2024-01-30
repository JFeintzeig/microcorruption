#!/bin/bash

mv ~/Downloads/memory\ \(1\).bin mem.bin

# requires hash dump starting at 0x8200
# xxd -p -s 33280 -l 5120 -c 32 mem.bin > hashes.txt
xxd -p -s 33280 -l 5120 -c 4 mem.bin > hashes.txt

# the above is useful for debugger mode. for real mode
# we can only use what we print to screen, which we
# hand tabulate and hardcode into main.go as printed_hashes.txt
go run main.go
