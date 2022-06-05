package bloomfilter

type Bits []byte

type HashFunc[T any] func(n, k uint, v T) uint

func update[T any](b Bits, k uint, f HashFunc[T], v T) {
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
