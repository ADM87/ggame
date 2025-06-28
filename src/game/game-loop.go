package game

import (
	"github.com/hajimehoshi/ebiten"
)

type Loop interface {
	Update(*ebiten.Image) error
}

type gameloop struct {
}

func NewGameLoop() Loop {
	return &gameloop{}
}

func (g *gameloop) Update(image *ebiten.Image) error {
	return nil
}
