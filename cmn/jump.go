package cmn

import (
	"hash"
	"io"
	"hash/fnv"
)

var (
	FNV    hash.Hash64 = fnv.New64a()
	BakNum             = 2
)

func jumpHash(key uint64, buckets int) int {
	var b, j int64
	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int(b)
}

func hashInt(key uint64, buckets int, length int) []int {
	res := make([]int, 0, length)

	if buckets <= 0 {
		buckets = 1
	}

	for i := 0; i < length; {
		h := jumpHash(key, buckets)
		for j := 0; j < len(res); j++ {
			if res[j] == h {
				h = -1
				break
			}
			if res[j] > h {
				res[j], h = h, res[j]
			}
		}
		if h >= 0 {
			res = append(res, h)
			i++
		}
		key++

	}
	return res
}

func hashString(key string, buckets int, length int, keyHasher hash.Hash64) []int {
	keyHasher.Reset()

	_, err := io.WriteString(keyHasher, key)
	if err != nil {
		panic(err)
	}
	return hashInt(keyHasher.Sum64(), buckets, length)
}
