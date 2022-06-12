package avatar

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar/encoder"
	"github.com/AtilioBoher/integrador-avater-me/avatar/images"
)

// encoderInter is someone who can encode information.
type encoderInter interface {
	EncodeInfo(strInfo string) (encodedInfo []byte, err error)
}

// imageGenInter is someone who can make the identicons.
type imageGenInter interface {
	BuildAndSaveImage(encodedInfo []byte, filePath string) error
}

// Service contains functionalities related to avatar generation.
type avatarGenerator struct {
	Encoder   encoderInter
	Generator imageGenInter
}

// Info can contain the information related to the generation of an avatar.
// StrInfo contains the string that is going to be hashed to ganerate the avatar,
// and FilePath is the path where the avatar will be stored (the image format must be .png)
type Info struct {
	StrInfo  string
	FilePath string
}

// Returns the pointer to an avatar generator
func GimmeAnAvatarGenerator() *avatarGenerator {
	a := avatarGenerator{
		Encoder:   encoder.GimmeAnEncoder(),
		Generator: images.GimmeAnImageGenerator(),
	}
	return &a
}

// Check if the info struct is empty and returns an error if that is the case
func (i *Info) isInfoEmpty() error {
	empty := Info{}
	if *i == empty {
		return fmt.Errorf("error: the Info struct supplied is empty, please insert information before attempting to generate an avatar")
	}
	return nil
}

// Generates and store an avatar, needs a filled Info struct
func (a *avatarGenerator) GenerateAndSaveAvatar(info Info) error {
	if err := info.isInfoEmpty(); err != nil {
		return err
	}

	hash, err := a.Encoder.EncodeInfo(info.StrInfo)
	if err != nil {
		return err
	}

	err = a.Generator.BuildAndSaveImage(hash, info.FilePath)
	if err != nil {
		return err
	}

	return nil
}
