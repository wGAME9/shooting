package object

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type level struct {
	water Object

	duckImage  *ebiten.Image
	ducks      []Object
	maxOfDucks int
}

func NewLevel(
	waterImage,
	duckImage *ebiten.Image,
	maxOfDucks int,
) Object {
	return &level{
		water:      NewWater(waterImage),
		duckImage:  duckImage,
		ducks:      make([]Object, 0, maxOfDucks),
		maxOfDucks: maxOfDucks,
	}
}

func (l *level) Update(screen *ebiten.Image, tick uint) error {
	if err := l.water.Update(screen, tick); err != nil {
		return err
	}

	if len(l.ducks) < l.maxOfDucks {
		screenHeight := screen.Bounds().Size().Y
		duckImageHeight := l.duckImage.Bounds().Size().Y
		y := screenHeight - duckImageHeight*3/2
		// every second there's 40% possibilities to
		// generate a missing duck
		if tick%60 == 0 && rand.Float64() < 0.4 {
			l.ducks = append(l.ducks, NewDuck(l.duckImage, float64(y)))
		}
	}

	n := 0
	for _, duck := range l.ducks {
		if err := duck.Update(screen, tick); err != nil {
			return err
		}

		if duck.IsOnScreen() {
			l.ducks[n] = duck
			n++
		}
	}

	l.ducks = l.ducks[:n]

	return nil
}

func (l *level) Draw(screen *ebiten.Image) {
	for _, duck := range l.ducks {
		duck.Draw(screen)
	}

	l.water.Draw(screen)
}

func (l *level) IsOnScreen() bool {
	return l.water.IsOnScreen()
}
