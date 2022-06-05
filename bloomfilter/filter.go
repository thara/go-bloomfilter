package bloomfilter

type HashFunc[T any] func(n, k uint, v T) uint

type BloomFilter[T any] struct {
	k uint
	f HashFunc[T]
}

func New[T any](k uint, f HashFunc[T]) BloomFilter[T] {
	return BloomFilter[T]{
		k: k,
		f: f,
	}
}

func (f *BloomFilter[T]) Set(b []byte, v T)       { update(b, f.k, f.f, v) }
func (f *BloomFilter[T]) Test(b []byte, v T) bool { return filter(b, f.k, f.f, v) }

func update[T any](b []byte, k uint, f HashFunc[T], v T) {
	n := uint(len(b))
	for i := uint(0); i < k; i++ {
		h := f(n, i, v)
		col := h % 8
		row := h / 8
		b[row] |= 1 << (7 - col)
	}
}

func filter[T any](b []byte, k uint, f HashFunc[T], v T) bool {
	n := uint(len(b))
	for i := uint(0); i < k; i++ {
		h := f(n, i, v)
		col := h % 8
		row := h / 8
		if (b[row]>>(7-col))&1 == 0 {
			return false
		}
	}
	return true
}
