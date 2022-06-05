package testutil

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash"
	"io"
)

func MyHash(k uint, v string) uint {
	//NOTE: k < K
	var h hash.Hash
	switch k {
	case 0:
		h = sha1.New()
	case 1:
		h = sha256.New()
	case 2:
		h = sha512.New()
	case 3:
		h = md5.New()
	}
	_, err := io.WriteString(h, v)
	if err != nil {
		panic(err) // need to error handing
	}
	r := h.Sum(nil)
	buf := bytes.NewBuffer(r)
	x, _ := binary.ReadUvarint(buf)
	return uint(x)
}
