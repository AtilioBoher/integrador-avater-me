/*
Package images provides a structure type called icon, which implements the imageGen
interface.
*/
package images

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// this constants can, but are not meant to be changed. Do so at your discretion
const (
	imageSize = 300 // generate image of size (imageSize x ImageSize) pixels
	blockSize = 10  // generate icon of (IconSize x blockSize) blocks
	max8bits  = 255 // Maximum 8 bit number in decimal
	// Minimun Hash Length, the hash can't be smaller than minHashLen
	minHashLen = 20
	// Extension used for the generated image
	fileExt = ".png"

	pixelsPerBlock = imageSize / blockSize // number of pixels per block
)

// Magic values
var colorWhite = color.NRGBA{max8bits, max8bits, max8bits, max8bits} // color white
var colorRed = color.NRGBA{max8bits, 0, 0, max8bits}                 // color red
var specialHash = []byte{121, 63, 78, 172, 17, 75, 25, 117, 193, 98, 8, 224, 206, 182, 251, 161, 39,
	198, 225, 193}

type icon struct {
	grid           [][]color.Color
	blockColor1    color.Color
	blockColor2    color.Color
	img            *image.NRGBA
	imageSize      int
	blockSize      int
	pixelsPerBlock int
}

// NewImageGenerator returns a pointer to an image generator which implements the ImageGen
// interface
func NewImageGenerator() *icon {
	ic := icon{
		grid:           [][]color.Color{},
		blockColor1:    nil,
		blockColor2:    nil,
		img:            &image.NRGBA{},
		imageSize:      imageSize,
		blockSize:      blockSize,
		pixelsPerBlock: pixelsPerBlock,
	}
	// generates a blank image where the identicon will be inserted, block by block
	for i := 0; i < ic.imageSize; i++ {
		var y []color.Color
		for j := 0; j < ic.imageSize; j++ {
			y = append(y, colorWhite)
		}
		ic.grid = append(ic.grid, y)
	}
	return &ic
}

// makeGrid takes a 20 bytes long hash, or larger, and generates the grid of the identicon from it.
// The identicon generated is vertically symetric. If the hash is larger than 20 bytes, excess bytes
// will be ignored.
func (ic *icon) makeGrid(hash []byte) error {
	if len(hash) < minHashLen {
		return fmt.Errorf("error: hash has to be at least 20 bytes long")
	}

	if isSpecialPerson(hash) { // this is an easter egg, it generates a heart identicon
		// if the hash is generated with the name of that special person
		generateHeart(ic)
		return nil
	}

	ic.blockColor1 = color.NRGBA{hash[minHashLen-6], hash[minHashLen-5], hash[minHashLen-4], max8bits}
	ic.blockColor2 = color.NRGBA{hash[minHashLen-3], hash[minHashLen-2], hash[minHashLen-1], max8bits}
	var aux byte
	for x := 0; x < ic.blockSize/2; x++ {
		for y := 0; y < ic.blockSize; y++ {
			aux = hash[x] + hash[y]
			switch {
			case aux < max8bits/3:
				ic.putColorAtBlock(x, y, ic.blockColor2)
				//this is the vertically mirrored side
				ic.putColorAtBlock(ic.blockSize-1-x, y, ic.blockColor2)
			case aux < max8bits*2/3:
				ic.putColorAtBlock(x, y, ic.blockColor1)
				//this is the vertically mirrored side
				ic.putColorAtBlock(ic.blockSize-1-x, y, ic.blockColor1)
			}
		}
	}
	return nil
}

// generateImage generates an image from the grid in the icon Struct
func (ic *icon) generateImage() {
	xlen, ylen := len(ic.grid), len(ic.grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	ic.img = image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			ic.img.Set(x, y, ic.grid[x][y])
		}
	}
}

// saveImage saves the image as a file in the specified filePath, the format has to be .png or
// an error will occur
func (ic *icon) saveImage(filePath string) error {
	if fileExtNotPng(filePath) {
		return fmt.Errorf("error: file extension specified in the path is not in .png format")
	}
	imgFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	png.Encode(imgFile, ic.img.SubImage(ic.img.Rect))
	return nil
}

// fileExtNotPng returns true if the extensión of filePath is not .png
func fileExtNotPng(filePath string) bool {
	return string(filePath[len(filePath)-len(fileExt):]) != fileExt
}

// putColorAtBlock put color at the block (x,y) specified.
func (ic *icon) putColorAtBlock(x, y int, blockColor color.Color) {
	for i := x * ic.pixelsPerBlock; i < ((x + 1) * ic.pixelsPerBlock); i++ {
		for j := y * ic.pixelsPerBlock; j < ((y + 1) * ic.pixelsPerBlock); j++ {
			ic.grid[i][j] = blockColor
		}
	}
}

// generateHeart is an easter egg, it generates a heart identicon
func generateHeart(ic *icon) {
	// top and bottom are magic values, but since they're not meant to be changed, are left here
	top := []int{2, 1, 1, 2, 3}
	bottom := []int{6, 7, 8, 9, 10}
	for x := 0; x < ic.blockSize/2; x++ {
		for y := top[x]; y < bottom[x]; y++ {
			ic.putColorAtBlock(x, y, colorRed)
			ic.putColorAtBlock(ic.blockSize-1-x, y, colorRed) //this is the vertically mirrored side
		}
	}
}

// isSpecialPerson returns true if the hash match specialHash
func isSpecialPerson(hash []byte) bool {
	for i := range hash {
		if hash[i] != specialHash[i] {
			return false
		}
	}
	return true
}

// BuildAndSaveImage generates an identicon with the hash "encodedInfo" and store it in the path supplied.
func (ic *icon) BuildAndSaveImage(encodedInfo []byte, filePath string) error {
	err := ic.makeGrid(encodedInfo)
	if err != nil {
		return err
	}
	ic.generateImage()
	err = ic.saveImage(filePath)
	if err != nil {
		return err
	}
	return nil
}
