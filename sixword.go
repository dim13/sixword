// Package sixword implements Six Word Format as in RFC 2289
package sixword

import "strings"

// Encode 64 bit unsigned integer value to Six-Word string
func Encode(n uint64) string {
	s := []string{
		dict[n>>53&0x7ff],
		dict[n>>42&0x7ff],
		dict[n>>31&0x7ff],
		dict[n>>20&0x7ff],
		dict[n>>9&0x7ff],
		dict[checksum(n)|n<<2&0x7fc],
	}
	return strings.Join(s, " ")
}

// Decode Six-Word string to 64 bit unsigned integer value
func Decode(s string) (uint64, bool) {
	var x [6]uint64
	words := strings.Fields(s)
	if len(words) != len(x) {
		return 0, false
	}
	for i, v := range words {
		n, ok := index[v]
		if !ok {
			return 0, false
		}
		x[i] = uint64(n)
	}
	n := x[0]<<53 | x[1]<<42 | x[2]<<31 | x[3]<<20 | x[4]<<9 | x[5]>>2
	return n, x[5]&3 == checksum(n)
}

func checksum(n uint64) uint64 {
	var sum uint64
	for i := range 32 {
		sum += n >> (i << 1) & 3
	}
	return sum & 3
}
