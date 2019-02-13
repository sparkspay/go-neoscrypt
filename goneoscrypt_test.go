package goneoscrypt

import (
	"bytes"
	"testing"
)

func TestNeoScryptHash(t *testing.T) {
	neoscryptHash := NeoscryptHash([]byte("00000000000000000000000000000000000000000000000000000000000000000000000000000000"), 0)
	exp := [32]byte{
		110, 239, 215, 250, 64, 118, 73, 49, 66, 86, 57, 94, 63, 230, 30, 131, 61, 28, 81, 148, 226, 99, 103, 87, 251, 112, 186, 1, 144, 167, 192, 147}
	if bytes.Compare(neoscryptHash[:], exp[:]) != 0 {
		t.Error("Hashing produced", neoscryptHash, "expected", exp)
	}
}
