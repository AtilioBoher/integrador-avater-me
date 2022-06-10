package avatar

import (
	"github.com/AtilioBoher/integrador-avater-me/avatar/encoder"
	"github.com/AtilioBoher/integrador-avater-me/avatar/images"
)

// encoderInter is someone who can encode information.
type encoderInter interface {
	EncodeInfo(strInfo string) (encodedInfo []byte, err error)
}

// imageInter is someone who can make images.
type imageGenInter interface {
	BuildAndSaveImage(encodedInfo []byte) error
}

// Service contains functionalities related to avatar generation.
type avatarGenerator struct {
	Encoder   encoderInter
	Generator imageGenInter
}

// Information contains information
type Info struct {
	Email string
}

func GimmeAnAvatarGenerator() *avatarGenerator {
	a := avatarGenerator{
		Encoder:   encoder.GimmeAnEncoder(),
		Generator: images.GimmeAnImageGenerator(),
	}
	return &a
}

func (a *avatarGenerator) GenerateAndSaveAvatar(info Info) error {
	// here will be all the logic
	hash, err := a.Encoder.EncodeInfo(info.Email)
	if err != nil {
		return err
	}
	err = a.Generator.BuildAndSaveImage(hash)
	if err != nil {
		return err
	}
	return nil
}
