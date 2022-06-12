package main

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar"
)

func main() {
	info := avatar.Info{
		StrInfo:  "atilio",
		FilePath: "identicon.png",
	}

	a := avatar.GimmeAnAvatarGenerator()
	err := a.GenerateAndSaveAvatar(info)
	if err != nil {
		fmt.Println(err)
	}

}
