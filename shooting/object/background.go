package object

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type background struct {
	image         *ebiten.Image
	width, height int
}

func NewBackground(backgroundImage *ebiten.Image) Object {
	size := backgroundImage.Bounds().Size()
	width, height := size.X, size.Y

	return &background{
		image:  backgroundImage,
		width:  width,
		height: height,
	}
}

func (bg *background) Update(screen *ebiten.Image, tick uint) error {
	return nil
}

func (bg *background) Draw(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()
	screenWidth, screenHeight := screenSize.X, screenSize.Y

	for i := range screenWidth / bg.width {
		for j := range screenHeight / bg.height {
			op := &ebiten.DrawImageOptions{}
			x, y := i*bg.width, j*bg.height
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(bg.image, op)
		}
	}
}

func (bg *background) IsOnScreen() bool {
	return true
}
