package main

import (
	"fmt"
	"strconv"
)

// This file is for the mining code.
// Note that "targetBits" for this assignment, at least initially, is 33.
// This could change during the assignment duration!  I will post if it does.

// Mine mines a block by varying the nonce until the hash has targetBits 0s in
// the beginning.  Could take forever if targetBits is too high.
// Modifies a block in place by using a pointer receiver.

func (self Block) Mine() {
	nonce := 1
	self.Nonce = string(nonce)
	h := self.Hash()
	verified := Verify(h, targetBits)

	for !verified && !reset {
		nonce += 1
		self.Nonce = strconv.Itoa(nonce)
		//fmt.Println("\nnonce %s", nonce)
		h = self.Hash()
		verified = Verify(h, targetBits)
	}
	if verified {
		SendBlockToServer(self)
		fmt.Println("Blocked mined, nonce \n%c", nonce)
	}

}

func Verify(blh Hash, t int) bool {
	//	fmt.Printf("\n%x", blh)
	var v bool
	for i := 0; i < t; i++ {
		if blh[i/8]>>uint(7-(i%8))&0x01 == 0 {
			v = true
		} else {
			v = false
		}
		if v == false {
			break
		}
		//		fmt.Printf("\n%x", i)
	}
	return v

}
