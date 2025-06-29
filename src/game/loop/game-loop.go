package loop

import (
	"github.com/ADM87/ggame/src/keyboard"
)

type Loop interface {
	Update() error
}

type gameloop struct {
}

func NewGameLoop() Loop {
	return &gameloop{}
}

func (g *gameloop) Update() error {
	if err := keyboard.Ping(); err != nil {
		return err
	}
	return nil
}
