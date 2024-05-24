package shooting

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 768
)

type game struct {
	backgroundImage *ebiten.Image
	background      background

	tick int
}

func NewGame() ebiten.Game {
	return &game{
		backgroundImage: ebiten.NewImage(screenWidth, screenHeight),
		background:      background{},
	}
}

func (g *game) Update() error {
	g.tick++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.background.Draw(g.backgroundImage, g.tick)
	screen.DrawImage(g.backgroundImage, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
