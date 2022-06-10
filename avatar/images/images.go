package images

func IsImagesWorking() bool {
	return true
}

type imageGen struct {
}

func GimmeAnImageGenerator() *imageGen { // returns an uninitialized imageGenerator with zero values
	i := imageGen{}
	return &i
}

func (i *imageGen) BuildAndSaveImage(encodedInfo []byte) error {
	return nil
}
