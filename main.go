package main

import (
    "fmt"
    "github.com/AtilioBoher/integrador-avater-me/avatar/encoder"
    "github.com/AtilioBoher/integrador-avater-me/avatar/images"
)

func main() {
	fmt.Println("hello world")

    if encoder.IsEncoderWorking() {
        fmt.Println("package encoder seems to be working")
    }

    if images.IsImagesWorking() {
        fmt.Println("package images seems to be working")
    }
}