package encoder

import (
	"crypto/sha256"
	"fmt"
)

func IsEncoderWorking() bool {
	return true
}

type encoder struct {
}

// returns an encoder
func GimmeAnEncoder() *encoder { 
	e := encoder{}
	return &e
}

func (e *encoder) EncodeInfo(strInfo string) (encodedInfo []byte, err error) {
	// in case the input is empy the function will return an empty byte slice and an error
	if strInfo == "" {
		return []byte{}, fmt.Errorf("error: string supplied as input is empty")
	}
	h := sha256.Sum256([]byte(strInfo)) // returns a [32]byte hash
	return h[:], nil	// the [:] notation returns a slice which point to the underliying array
}
