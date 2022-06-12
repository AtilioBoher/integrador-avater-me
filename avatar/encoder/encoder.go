package encoder

import (
	"crypto/sha1"
	"fmt"
)

type encoder struct {
}

// Returns a pointer to an encoder which implements the encoderInter interface
func GimmeAnEncoder() *encoder {
	e := encoder{}
	return &e
}

// Returns the hash of the input string and an error.
// The hash is made with the sha1 algorithm, which returns a 20 bytes long hash
func (e *encoder) EncodeInfo(strInfo string) (encodedInfo []byte, err error) {
	// in case the input is empy the function will return an empty byte slice and an error
	if strInfo == "" {
		return []byte{}, fmt.Errorf("error: string supplied as input is empty")
	}
	h := sha1.Sum([]byte(strInfo)) // returns a [20]byte hash
	return h[:], nil               // the [:] notation returns a slice which point to the underliying array
}
