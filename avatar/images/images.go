package images

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// this constants can, but are not meant to be changed. Do so at your discretion
const imageSize = 300                       // generated image of size (imageSize x ImageSize) pixels
const iconSize = 10                         // generated icon of (IconSize x iconSize) blocks
const pixelsPerBlock = imageSize / iconSize // number of pixels per block

type icon struct {
	grid           [][]color.Color
	blockColor1    color.Color
	blockColor2    color.Color
	img            *image.NRGBA
	imageSize      int
	iconSize       int
	pixelsPerBlock int
}

// Returns a pointer to an image generator which implements the ImageGenInter interface
func GimmeAnImageGenerator() *icon {
	ic := icon{
		grid:           [][]color.Color{},
		blockColor1:    nil,
		blockColor2:    nil,
		img:            &image.NRGBA{},
		imageSize:      imageSize,
		iconSize:       iconSize,
		pixelsPerBlock: pixelsPerBlock,
	}
	// generates a blanck image where the identicon will be inserted, block by block
	whiteNRGBA := color.NRGBA{255, 255, 255, 255}
	for i := 0; i < ic.imageSize; i++ {
		var y []color.Color
		for j := 0; j < ic.imageSize; j++ {
			y = append(y, whiteNRGBA)
		}
		ic.grid = append(ic.grid, y)
	}
	return &ic
}

// Takes a 20 bytes long hash, or larger, and generates the grid of the identicon from it. The identicon generated is vertically symetric.
func (ic *icon) makeGrid(hash []byte) error {
	if len(hash) < 20 {
		return fmt.Errorf("error: hash has to be at least 20 bytes long")
	}

	if specialPerson(ic, hash) { // this is like an inside joke, it generates a heart identicon if the hash is generated with the name of that special person
		return nil
	}

	ic.blockColor1 = color.NRGBA{hash[14], hash[15], hash[16], 255}
	ic.blockColor2 = color.NRGBA{hash[17], hash[18], hash[19], 255}
	var aux byte
	for x := 0; x < ic.iconSize/2; x++ {
		for y := 0; y < ic.iconSize; y++ {
			aux = hash[x] + hash[y]
			switch {
			case aux < 85:
				ic.putColorAt(x, y, true)
				ic.putColorAt(ic.iconSize-1-x, y, true) //this is the vertically mirrored side
			case aux < 170:
				ic.putColorAt(x, y, false)
				ic.putColorAt(ic.iconSize-1-x, y, false) //this is the vertically mirrored side
			}
		}
	}
	return nil
}

// It generates an image from the grid in the icon Struct, you have to generate the grid
// with the method "makeGrid" before attempting to generate the image, or the result will be a blank image
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

// Saves the image as a file in the specified filePath, the format has to be .png or an error will occur
func (ic *icon) saveImage(filePath string) error {
	if string(filePath[len(filePath)-4:]) != ".png" {
		return fmt.Errorf("error: file path specified is not in .png format")
	}
	imgFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	png.Encode(imgFile, ic.img.SubImage(ic.img.Rect))
	return nil
}

// Put color at the point (x,y) specified. The variable col determine which of the two colors available will be use
func (ic *icon) putColorAt(x, y int, col bool) {
	auxBlockColor := ic.blockColor1
	if col {
		auxBlockColor = ic.blockColor2
	}
	for i := x * ic.pixelsPerBlock; i < ((x + 1) * ic.pixelsPerBlock); i++ {
		for j := y * ic.pixelsPerBlock; j < ((y + 1) * ic.pixelsPerBlock); j++ {
			ic.grid[i][j] = auxBlockColor
		}
	}
}

// this is like an inside joke, it generates a heart identicon if the hash is generated with the name of that special person
func specialPerson(ic *icon, hash []byte) bool {
	specialHash := []byte{129, 249, 115, 24, 78, 33, 109, 185, 179, 234, 240, 10, 54, 12, 99, 156, 108, 24, 243, 171}
	isSpecial := true
	// compare, if they don't match returns to normal procedure
	for i := range hash {
		if hash[i] != specialHash[i] {
			isSpecial = false
		}
	}
	if !isSpecial {
		return false // is not Special, just return and keep going with the normal procedure
	}
	ic.blockColor2 = color.NRGBA{225, 0, 0, 255}
	// procede to generate a heart
	top := []int{2, 1, 1, 2, 3}
	bottom := []int{6, 7, 8, 9,10}
	for x := 0; x < ic.iconSize/2; x++ {
		for y := top[x]; y < bottom[x]; y++ {
			ic.putColorAt(x, y, true)
			ic.putColorAt(ic.iconSize-1-x, y, true) //this is the vertically mirrored side
		}
	}
	return true
}

// It generates an identicon with the hash "encodedInfo" and store it in the path supplied.
// Take into account that this implementation is ment to work with a 20 bytes long hash or longer
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
