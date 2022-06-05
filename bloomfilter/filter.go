package bloomfilter

// HashFunc is a hash function that be expected to generate a uniform random distribution.
//
// `T` is the type of target to hashing.
//
// `k` is an index in evaluations of multiple hash functions.
// `v` is a target of hashing.
type HashFunc[T any] func(k uint, v T) uint

// BloomFilter implements bloom filter algorithm.
//
// This represents bloom filter's algorithmic parameters.
// In other words, it doesn't have a bit array that constructed by the algorithm.
type BloomFilter[T any] struct {
	k uint
	f HashFunc[T]
}

// New returns new BloomFilter for `T` as hashing target.
// `k` is number of hash functions.
// `f` is base of hash function parameterized by `k`.
func New[T any](k uint, f HashFunc[T]) BloomFilter[T] {
	return BloomFilter[T]{
		k: k,
		f: f,
	}
}

// Set updates `b` as bit array by hashing for `v`.
func (f *BloomFilter[T]) Set(b []byte, v T) {
	n := uint(len(b)) * 8
	for i := uint(0); i < f.k; i++ {
		h := f.f(i, v) % n
		col := h % 8
		row := h / 8
		b[row] |= 1 << (7 - col)
	}
}

// Test returns true if `b` as bit array contains the result of hashing for `v`.
func (f *BloomFilter[T]) Test(b []byte, v T) bool {
	n := uint(len(b)) * 8
	for i := uint(0); i < f.k; i++ {
		h := f.f(i, v) % n
		col := h % 8
		row := h / 8
		if (b[row]>>(7-col))&1 == 0 {
			return false
		}
	}
	return true
}
