package shooting

import (
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
}
