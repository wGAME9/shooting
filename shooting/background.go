package shooting

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type background struct{}

func (b background) Draw(backgroundImage *ebiten.Image) {
	drawBackgroundGreen(backgroundImage)
	drawBackgroundWood(backgroundImage)
	drawCurtain(backgroundImage)
	drawCurtainStraight(backgroundImage)
}

func drawBackgroundGreen(backgroundImage *ebiten.Image) {
	imageWidth := 256
	imageHeight := 256
	numI := screenWidth / imageWidth
	numJ := screenHeight / imageHeight
	for i := range numI {
		for j := range numJ {
			op := &ebiten.DrawImageOptions{}
			x := i * imageWidth
			y := j * imageHeight
			op.GeoM.Translate(float64(x), float64(y))
			backgroundImage.DrawImage(backgroundGreenImage, op)
		}
	}
}

func drawBackgroundWood(backgroundImage *ebiten.Image) {
	imageWidth := 256
	imageHeight := 256
	numI := screenWidth / imageWidth
	fixedY := screenHeight - imageHeight/2
	for i := range numI {
		op := &ebiten.DrawImageOptions{}
		x := i * imageWidth
		op.GeoM.Translate(float64(x), float64(fixedY))
		backgroundImage.DrawImage(backgroundWoodImage, op)
	}
}

func drawCurtainStraight(backgroundImage *ebiten.Image) {
	imageWidth := 256

	numI := screenWidth / imageWidth
	for i := range numI {
		op := &ebiten.DrawImageOptions{}
		x := i * imageWidth
		op.GeoM.Translate(float64(x), 0)
		backgroundImage.DrawImage(curtainStraightImage, op)
	}
}

func drawCurtain(backgroundImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	backgroundImage.DrawImage(curtainImage, op)

	op.GeoM.Reset()
	op.GeoM.Scale(-1, 1) // Flip horizontally
	op.GeoM.Translate(float64(screenWidth), 0)
	backgroundImage.DrawImage(curtainImage, op)
}
