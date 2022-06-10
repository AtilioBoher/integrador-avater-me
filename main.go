package main

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar"
)

func main() {
	fmt.Println("hello world")

	info := avatar.Info{Email: "atilioboher@gmail.com"}
	fmt.Println(info)

	a := avatar.GimmeAnAvatarGenerator()
	fmt.Println(a)

}
