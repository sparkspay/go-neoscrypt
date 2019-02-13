package goneoscrypt

// #include "neoscrypt.h"
import "C"

// NeoscryptHash the provided data, returning a slice of the [32]byte containing the resulting hash.
func NeoscryptHash(data []byte, profile uint) [32]byte {
	var retbuf [32]byte
	C.neoscrypt((*C.uchar)(&data[0]), (*C.uchar)(&retbuf[0]), C.uint(profile))
	return retbuf
}
