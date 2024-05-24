package shooting

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type background struct {
	spriteSheet objectSpriteSheet
}

func (b background) Draw(backgroundImage *ebiten.Image, tick int) {
	drawBackgroundGreen(backgroundImage)
	drawDuckImage(backgroundImage, b.spriteSheet, tick)
	drawWater(backgroundImage, tick)
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

var (
	waterDirection int = 1
	waterStartingX int = -132
)

func drawWater(backgroundImage *ebiten.Image, tick int) {
	imageWidth := 132
	imageHeight := 224
	// draw 3 more to cover the screen
	// since the water is moving
	numI := screenWidth/imageWidth + 3
	y := screenHeight - imageHeight

	if int(tick)%60 == 0 {
		waterDirection *= -1
	}
	waterStartingX += waterDirection

	for i := range numI {
		op := &ebiten.DrawImageOptions{}
		realX := i*imageWidth + waterStartingX
		op.GeoM.Translate(float64(realX), float64(y))
		backgroundImage.DrawImage(waterImage, op)
	}
}

var (
	duckStartingX  = -114
	duckStartingY  = screenHeight - 420
	duckYDirection = 1
)

func drawDuckImage(backgroundImage *ebiten.Image, spriteSheet objectSpriteSheet, tick int) {
	if duckStartingX == screenWidth {
		return
	}

	realDuckImage := getDuckImage(spriteSheet)

	if tick%2 == 0 {
		duckStartingX += 2
	}

	if tick%60 == 0 {
		duckYDirection *= -1
	}
	duckStartingY += duckYDirection

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(duckStartingX), float64(duckStartingY))
	backgroundImage.DrawImage(realDuckImage, op)
}

func getDuckImage(spriteSheet objectSpriteSheet) *ebiten.Image {
	duckImag, err := spriteSheet.GetImage("duck_outline_target_white.png")
	if err != nil {
		log.Fatal(err)
	}
	duckImageSize := duckImag.Bounds().Size()
	duckWidth, duckHeight := duckImageSize.X, duckImageSize.Y

	stickImage, err := spriteSheet.GetImage("stick_woodFixed_outline.png")
	if err != nil {
		log.Fatal(err)
	}
	stickImageSize := stickImage.Bounds().Size()
	stickWidth, stickHeight := stickImageSize.X, stickImageSize.Y

	squareWidth := duckWidth
	if stickWidth > duckWidth {
		squareWidth = stickWidth
	}
	squareHeight := duckHeight + stickHeight
	square := ebiten.NewImage(squareWidth, squareHeight)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	square.DrawImage(duckImag, op)

	stickStartingX := (squareWidth - stickWidth) / 2
	op.GeoM.Reset()
	op.GeoM.Translate(float64(stickStartingX), float64(duckHeight))
	square.DrawImage(stickImage, op)

	return square
}
