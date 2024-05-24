package object

import "github.com/hajimehoshi/ebiten/v2"

type desk struct {
	image *ebiten.Image
	width     int
	height    int
}

func NewDesk(deskImage *ebiten.Image) Object {
	size := deskImage.Bounds().Size()
	width, height := size.X, size.Y

	return &desk{
		image: deskImage,
		width:     width,
		height:    height,
	}
}

func (d *desk) Update(screen *ebiten.Image, tick uint) error {
	return nil
}

func (d *desk) Draw(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()
	screenWidth, screenHeight := screenSize.X, screenSize.Y

	y := screenHeight - d.height/2
	for i := range screenWidth / d.width {
		op := &ebiten.DrawImageOptions{}
		x := i * d.width
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(d.image, op)
	}
}

func (d *desk) IsOnScreen() bool {
	return true
}
