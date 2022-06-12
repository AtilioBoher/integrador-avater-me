package main

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar"
)

func main() {
	info := avatar.Info{Email: "Samantha",
		Ip: "127.0.0.1"}

	a := avatar.GimmeAnAvatarGenerator()
	err := a.GenerateAndSaveAvatar(info)
	if err != nil {
		fmt.Println(err)
	}

}
