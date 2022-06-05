package avatar

// cryptoEncoder is someone who can encode information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

// imageGenerator is someone who can make images.
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

// Service contains functionalities related to avatar generation.
type Service struct {
	encoder   cryptoEncoder
	generator imageGenerator
}

// Information contains information (?)
type Information struct {
	// here goes all the information you want to encode
}

func (s *Service) GenerateAndSaveAvatar(information Information) error {
	// here will be all the logic
	return nil
}
