package object

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ducksXSpeed     = 1.5 // horizontal speed
	ducksYSpeed     = 0.6 // vertical speed
	ducksMaxOffsetY = 16  // max vertical movement for animation
)

type duck struct {
	image  *ebiten.Image
	width  int
	height int

	startingX  float64
	startingY  float64
	dy         float64
	yDirection int

	isOnScreen bool
}

func newDuck(duckImage *ebiten.Image, initialY float64) *duck {
	size := duckImage.Bounds().Size()
	width, height := size.X, size.Y

	return &duck{
		image:      duckImage,
		width:      width,
		height:     height,
		startingX:  -float64(width),
		startingY:  initialY,
		yDirection: 1,
		isOnScreen: true,
	}
}

func (d *duck) Update(screen *ebiten.Image, tick uint) error {
	screenWidth := screen.Bounds().Size().X
	if int(d.startingX) >= screenWidth {
		d.isOnScreen = false
		return nil
	}

	d.startingX += ducksXSpeed

	// calculate the vertical direction and offset (for animation)
	if ducksMaxOffsetY-math.Abs(d.dy) < 0 {
		d.yDirection *= -1
	}
	d.dy += float64(d.yDirection) * ducksYSpeed

	return nil
}

func (d *duck) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.startingX, d.startingY+d.dy)
	screen.DrawImage(d.image, op)
}

func (d *duck) IsOnScreen() bool {
	return d.isOnScreen
}
