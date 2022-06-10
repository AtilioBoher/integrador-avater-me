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

// imageInter is someone who can make images.
type imageGenInter interface {
	BuildAndSaveImage(encodedInfo []byte) error
}

// Service contains functionalities related to avatar generation.
type avatarGenerator struct {
	Encoder   encoderInter
	Generator imageGenInter
}

// Info can contain either an email or an IP or both,
// but it can't be left empty or an error will occur when used
type Info struct {
	Email string
	Ip    string
}

func GimmeAnAvatarGenerator() *avatarGenerator {
	a := avatarGenerator{
		Encoder:   encoder.GimmeAnEncoder(),
		Generator: images.GimmeAnImageGenerator(),
	}
	return &a
}

func (i *Info) isInfoEmpty() error {
	empty := Info{}
	if *i == empty {
		return fmt.Errorf("error: the Info struct supplied is empty, please insert and email or an ip address")
	}
	return nil
}

func (a *avatarGenerator) GenerateAndSaveAvatar(info Info) error {
	// here will be all the logic

	if err := info.isInfoEmpty(); err != nil {
		return err
	}

	// if in case both are supplied, email will be used
	inputInfo := info.Ip
	if info.Email != "" {
		inputInfo = info.Email
	}

	hash, err := a.Encoder.EncodeInfo(inputInfo)
	if err != nil {
		return err
	}

	fmt.Println(hash)

	/* err = a.Generator.BuildAndSaveImage(hash)
	if err != nil {
		return err
	} */

	return nil
}
