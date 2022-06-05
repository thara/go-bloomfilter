package bloomfilter

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"io"
	"testing"
)

func Test_update(t *testing.T) {
	id := func(n, k uint, v int) uint {
		return uint(v)
	}
	tests := []struct {
		name string
		b    []byte
		v    int
		want []byte
	}{
		{
			name: "0",
			b:    []byte{0b00000000},
			v:    0,
			want: []byte{0b10000000},
		},
		{
			name: "1",
			b:    []byte{0b00000000},
			v:    1,
			want: []byte{0b01000000},
		},
		{
			name: "4",
			b:    []byte{0b00000001},
			v:    4,
			want: []byte{0b00001001},
		},
		{
			name: "7",
			b:    []byte{0b10010000},
			v:    7,
			want: []byte{0b10010001},
		},
		{
			name: "8",
			b:    []byte{0b10010000, 0b00010000},
			v:    8,
			want: []byte{0b10010000, 0b10010000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			update(tt.b, 1, id, tt.v)

			for i, got := range tt.b {
				if got != tt.want[i] {
					t.Fatalf("want %v, but %v:", tt.want[i], tt.b)
				}
			}
		})
	}

}

func Test_filter(t *testing.T) {
	id := func(n, k uint, v int) uint {
		return uint(v)
	}

	tests := []struct {
		name string
		b    []byte
		v    int
		want bool
	}{
		{
			name: "1 - matched",
			b:    []byte{0b01000000},
			v:    1,
			want: true,
		},
		{
			name: "1 - unmatched",
			b:    []byte{0b00100000},
			v:    1,
			want: false,
		},
		{
			name: "4 - matched",
			b:    []byte{0b00001000},
			v:    4,
			want: true,
		},
		{
			name: "4 - unmatched",
			b:    []byte{0b00000100},
			v:    4,
			want: false,
		},
		{
			name: "7 - matched",
			b:    []byte{0b00001001},
			v:    7,
			want: true,
		},
		{
			name: "7 - unmatched",
			b:    []byte{0b00001000},
			v:    7,
			want: false,
		},
		{
			name: "10 - matched",
			b:    []byte{0b00000000, 0b00100000},
			v:    10,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filter(tt.b, 1, id, tt.v)
			if got != tt.want {
				t.Fatalf("want %v, but %v:", tt.want, got)
			}
		})
	}
}

func Test_bloomFilter(t *testing.T) {
	b := make([]byte, 256)

	h := func(n, k uint, v string) uint {
		h := sha1.New()
		_, err := io.WriteString(h, v)
		if err != nil {
			t.FailNow()
		}
		r := h.Sum(nil)
		buf := bytes.NewBuffer(r)
		x, _ := binary.ReadUvarint(buf)
		return uint(x) % n
	}

	update(b, 1, h, "abc")
	update(b, 1, h, "xxx")
	update(b, 1, h, "yyy")
	update(b, 1, h, "zzz")

	got := filter(b, 1, h, "abc")
	if got != true {
		t.Fatalf("want %v, but %v:", true, got)
	}

	tests := []struct {
		name string
		v    string
		want bool
	}{
		{
			name: "case1",
			v:    "abc",
			want: true,
		},
		{
			name: "case2",
			v:    "xxx",
			want: true,
		},
		{
			name: "case3",
			v:    "yyy",
			want: true,
		},
		{
			name: "case4",
			v:    "zzz",
			want: true,
		},
		{
			name: "case4",
			v:    "aaa",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filter(b, 1, h, tt.v)
			if got != tt.want {
				t.Fatalf("want %v, but %v:", tt.want, got)
			}
		})
	}
}
