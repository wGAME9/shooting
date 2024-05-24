package shooting

import (
	_ "embed"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// 256x256
	backgroundGreenImage *ebiten.Image
	// 256x256
	backgroundWoodImage *ebiten.Image
	// 256x80
	curtainStraightImage *ebiten.Image
	// 131x426
	curtainImage *ebiten.Image
	// 132x224
	waterImage *ebiten.Image

	objectsSpriteSheet *ebiten.Image
)

func init() {
	var err error

	backgroundGreenImage, _, err = ebitenutil.NewImageFromFile("assets/bg_green.png")
	if err != nil {
		log.Fatal(err)
	}

	backgroundWoodImage, _, err = ebitenutil.NewImageFromFile("assets/bg_wood.png")
	if err != nil {
		log.Fatal(err)
	}

	curtainStraightImage, _, err = ebitenutil.NewImageFromFile("assets/curtain_straight.png")
	if err != nil {
		log.Fatal(err)
	}

	curtainImage, _, err = ebitenutil.NewImageFromFile("assets/curtain.png")
	if err != nil {
		log.Fatal(err)
	}

	waterImage, _, err = ebitenutil.NewImageFromFile("assets/water1.png")
	if err != nil {
		log.Fatal(err)
	}

	objectsSpriteSheet, _, err = ebitenutil.NewImageFromFile("assets/objects.png")
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed objects.json
var objectsSpriteSheetBytes []byte

type objectSpriteSheet struct {
	Images []object `json:"images"`
}

func (s objectSpriteSheet) GetImage(name string) (*ebiten.Image, error) {
	for _, object := range s.Images {
		if object.Name == name {
			x, y := object.X, object.Y
			width, height := object.Width, object.Height

			rect := image.Rect(x, y, x+width, y+height)

			return objectsSpriteSheet.SubImage(rect).(*ebiten.Image), nil
		}
	}

	return nil, fmt.Errorf("image not found: [%s]", name)
}

type object struct {
	Name   string `json:"name"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
