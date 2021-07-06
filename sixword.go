// Package sixword implements Six Word Format as in RFC 2289

package sixword

import "strings"

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

func Decode(s string) (uint64, bool) {
	x, ok := parse(strings.Fields(s))
	if !ok || len(x) != 6 {
		return 0, false
	}
	n := x[0]<<53 | x[1]<<42 | x[2]<<31 | x[3]<<20 | x[4]<<9 | x[5]>>2
	return n, x[5]&3 == checksum(n)
}

func checksum(n uint64) uint64 {
	var sum uint64
	for i := 0; i < 32; i++ {
		sum += n >> (i << 1) & 3
	}
	return sum & 3
}

func parse(s []string) ([]uint64, bool) {
	x := make([]uint64, len(s))
	for i, v := range s {
		n, ok := index(v)
		if !ok {
			return nil, false
		}
		x[i] = n
	}
	return x, true
}

func index(s string) (uint64, bool) {
	s = strings.ToUpper(s)
	for i, v := range dict {
		if v == s {
			return uint64(i), true
		}
	}
	return 0, false
}
