package images

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

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
	whiteNRGBA := color.NRGBA{255, 255, 255, 255}
	for i := 0; i < imageSize; i++ {
		var y []color.Color
		for j := 0; j < imageSize; j++ {
			y = append(y, whiteNRGBA)
		}
		ic.grid = append(ic.grid, y)
	}
	return &ic
}

func (ic *icon) saveImage(filePath string) error {
	imgFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	png.Encode(imgFile, ic.img.SubImage(ic.img.Rect))
	return nil
}

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

func (ic *icon) makeGrid(hash []byte) {
	ic.blockColor1 = color.NRGBA{hash[14], hash[15], hash[16], 255}
	ic.blockColor2 = color.NRGBA{hash[17], hash[18], hash[19], 255}
	var aux byte
	for x := 0; x < ic.iconSize/2; x++ {
		for y := 0; y < ic.iconSize; y++ {
			aux = hash[x] + hash[y]
			switch {
			case aux < 85:
				ic.putColorAt(x, y, true)
				ic.putColorAt(ic.iconSize-1-x, y, true)
			case aux < 170:
				ic.putColorAt(x, y, false)
				ic.putColorAt(ic.iconSize-1-x, y, false)
			}
		}
	}
}

func (ic *icon) BuildAndSaveImage(encodedInfo []byte) error {
	ic.makeGrid(encodedInfo)
	ic.generateImage()
	ic.saveImage("random.png")
	return nil
}
