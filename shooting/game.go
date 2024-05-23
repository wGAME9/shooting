package shooting

import (
	"image"

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

const (
	imageSize = 16
	numFrames = 8
)

func (g *game) Draw(screen *ebiten.Image) {
	g.background.Draw(g.backgroundImage)

	speed := 10
	frameNum := int(g.tick/speed) % numFrames

	frameX := frameNum * imageSize
	rect := image.Rect(frameX, 0, frameX+imageSize, imageSize)
	subImage := coinSpriteSheet.SubImage(rect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4) // enlarge the coin (16x16 -> 64x64)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	g.backgroundImage.DrawImage(subImage, op)

	screen.DrawImage(g.backgroundImage, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
