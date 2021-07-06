package main

import (
	"crypto/rand"
	"encoding/binary"
)

// Source represents a source of random int64 values in the range [0, 1<<63).
type Source struct{}

// Seed is not used, required to satisfy rand.Source interface only.
func (s Source) Seed(seed int64) {}

// Int63 returns a non-negative random 63-bit integer as an int64.
func (s Source) Int63() int64 {
	return int64(s.Uint64() &^ uint64(1<<63))
}

// Uint64 returns a random 64-bit value as a uint64.
func (s Source) Uint64() uint64 {
	var v uint64
	if err := binary.Read(rand.Reader, binary.BigEndian, &v); err != nil {
		panic(err)
	}
	return v
}
