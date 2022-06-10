package encoder

func IsEncoderWorking() bool {
	return true
}

type encoder struct {
}

func GimmeAnEncoder() *encoder { // returns an uninitialized encoder with zero values
	e := encoder{}
	return &e
}

func (e *encoder) EncodeInfo(strInfo string) (encodedInfo []byte, err error) {
	return []byte{1, 2, 3}, nil
}
