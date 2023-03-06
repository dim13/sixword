package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/dim13/sixword"
)

func main() {
	n := rand.New(Source{}).Uint64()
	s := sixword.Encode(n)
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	fmt.Printf("%#0.16x %s\n", n, s)
}
