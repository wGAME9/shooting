package shooting

import (
	"encoding/json"
	"log"

	"github.com/wGAME9/shooting/shooting/object"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 768
)

type game struct {
	image *ebiten.Image

	objects []object.Object

	tick int
}

func NewGame() ebiten.Game {
	var spriteSheet objectSpriteSheet
	if err := json.Unmarshal(objectsSpriteSheetBytes, &spriteSheet); err != nil {
		log.Fatal(err)
	}

	return &game{
		image: ebiten.NewImage(screenWidth, screenHeight),
		objects: []object.Object{
			object.NewBackground(backgroundGreenImage),
			object.NewCurtain(
				curtainStraightImage,
				curtainImage,
			),
			object.NewLevel(waterImage, getDuckImage(spriteSheet), 5),
			object.NewDesk(backgroundWoodImage),
		},
	}
}

func (g *game) Update() error {
	g.tick++
	for _, obj := range g.objects {
		if err := obj.Update(g.image, uint(g.tick)); err != nil {
			return err
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for _, obj := range g.objects {
		obj.Draw(g.image)
	}

	screen.DrawImage(g.image, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
