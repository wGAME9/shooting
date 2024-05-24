package object

import "github.com/hajimehoshi/ebiten/v2"

type curtain struct {
	topCurtainImage      *ebiten.Image
	straightCurtainImage *ebiten.Image
}

func NewCurtain(
	topCurtainImage *ebiten.Image,
	straightCurtainImage *ebiten.Image,
) Object {
	return &curtain{
		topCurtainImage:      topCurtainImage,
		straightCurtainImage: straightCurtainImage,
	}
}

func (c *curtain) Update(screen *ebiten.Image, tick uint) error {
	return nil
}

func (c *curtain) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Size().X

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(c.straightCurtainImage, op)

	op.GeoM.Reset()
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(screenWidth), 0)
	screen.DrawImage(c.straightCurtainImage, op)

	topCurtainSize := c.topCurtainImage.Bounds().Size()
	topCurtainWidth := topCurtainSize.X
	numOfTopCurtains := screenWidth / topCurtainWidth
	for i := range numOfTopCurtains {
		op.GeoM.Reset()
		x := i * topCurtainWidth
		op.GeoM.Translate(float64(x), 0)
		screen.DrawImage(c.topCurtainImage, op)
	}
}

func (c *curtain) IsOnScreen() bool {
	return true
}
