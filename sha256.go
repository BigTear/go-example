package main

import (
	"crypto/sha256"
)

func main() {

	s := "1"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	println(s, bs)

}
