package main

import (
	"fmt"
	"math/rand"

	"github.com/dim13/sixword"
)

func main() {
	n := rand.New(Source{}).Uint64()
	s := sixword.Encode(n)
	fmt.Printf("%0.16x %s\n", n, s)
}
