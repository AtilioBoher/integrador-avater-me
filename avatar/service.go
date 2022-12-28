/*
Package avatar is capable of generating a user icon from a string with information from the user.
The information is hashed before is used to generate the icon, so that sensible information can
be used without security issues.
*/
package avatar

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar/encoder"
	"github.com/AtilioBoher/integrador-avater-me/avatar/images"
)

// infoEncoder is someone who can encode information.
type infoEncoder interface {
	EncodeInfo(strInfo string) (encodedInfo []byte, err error)
}

// imageGen is someone who can make identicons.
type imageGen interface {
	BuildAndSaveImage(encodedInfo []byte, filePath string) error
}

// avatarGenerator contains functionalities related to avatar generation.
type avatarGenerator struct {
	Encoder   infoEncoder
	Generator imageGen
}

// Info contains the information related to the generation of an avatar.
// StrInfo contains the string that is going to be hashed to ganerate the avatar,
// and FilePath is the path where the avatar will be stored (the image format must be .png).
type Info struct {
	StrInfo  string
	FilePath string
}

// NewAnAvatarGenerator returns the pointer to an avatarGenerator.
func NewAvatarGenerator() *avatarGenerator {
	return &avatarGenerator{
		Encoder:   encoder.NewEncoder(),
		Generator: images.NewImageGenerator(),
	}
}

// isEmpty checks if the info struct is empty and returns an error if that is the case.
func (i *Info) isEmpty() error {
	empty := Info{}
	if *i == empty {
		return fmt.Errorf("error: the Info struct supplied is empty, " +
			"please insert information before attempting to generate an avatar")
	}
	return nil
}

// GenerateAndSaveAvatar generates and store an avatar, needs a filled Info struct
func (a *avatarGenerator) GenerateAndSaveAvatar(info Info) error {
	if err := info.isEmpty(); err != nil {
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
