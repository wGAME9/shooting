package object

import "github.com/hajimehoshi/ebiten/v2"

type water struct {
	image     *ebiten.Image
	width     int
	height    int
	startingX int
	direction int
}

func NewWater(waterImage *ebiten.Image) Object {
	size := waterImage.Bounds().Size()
	width, height := size.X, size.Y

	return &water{
		image:     waterImage,
		width:     width,
		height:    height,
		startingX: -width,
		direction: 1,
	}
}

func (w *water) Update(screen *ebiten.Image, tick uint) error {
	if tick%60 == 0 {
		w.direction *= -1
	}
	w.startingX += w.direction

	return nil
}

func (w *water) Draw(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()
	screenWidth, screenHeight := screenSize.X, screenSize.Y

	y := screenHeight - w.height
	for i := range screenWidth/w.width + 2 {
		op := &ebiten.DrawImageOptions{}
		x := i*w.width + w.startingX
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(w.image, op)
	}
}

func (w *water) IsOnScreen() bool {
	return true
}
