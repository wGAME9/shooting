package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wGAME9/shooting/shooting"
)

func main() {
	game := shooting.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
