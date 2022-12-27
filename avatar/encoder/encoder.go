/*
Package encoder provides a structure type called encoder, which implements the infoEncoder
interface.
*/
package encoder

import (
	"crypto/sha1"
	"fmt"
)

type encoder struct {
}

// NewEncoder returns a new encoder.
func NewEncoder() *encoder {
	return &encoder{}
}

// EncodeInfo returns the hash of the input string and an error.
// The hash is made with the sha1 algorithm, which returns a 20 bytes long hash.
func (e *encoder) EncodeInfo(strInfo string) (encodedInfo []byte, err error) {
	if strInfo == "" {
		return []byte{}, fmt.Errorf("error: string supplied as input is empty")
	}
	h := sha1.Sum([]byte(strInfo)) // returns a [20]byte hash.
	return h[:], nil               // the [:] notation returns a slice which point to the
	// underliying array
}
