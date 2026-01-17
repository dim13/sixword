package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/dim13/sixword"
)

func main() {
	var n uint64
	binary.Read(rand.Reader, binary.BigEndian, &n)
	s := sixword.Encode(n)
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	fmt.Printf("%#0.16x %s\n", n, s)
}
