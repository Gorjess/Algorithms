package rabinkarp

import "github.com/cespare/xxhash"

func hash(src string) uint64 {
	return xxhash.Sum64String(src)
}

func XXHash(src string) uint64 {
	return hash(src)
}

func RK(src, pattern string) int {
	var (
		n = len(src)
		m = len(pattern)
	)

	hp := hash(pattern)
	subs := ""
	for i := 0; i < n-m+1; i++ {
		subs = src[i : i+m]
		if hash(subs) == hp {
			return i
		}
	}
	return -1
}
