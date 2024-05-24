package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Update(screen *ebiten.Image, tick uint) error
	Draw(screen *ebiten.Image)
	IsOnScreen() bool
}
