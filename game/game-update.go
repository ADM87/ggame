package game

import (
	"github.com/ADM87/ggame/keyboard"
)

func (g *gameshell) Update() error {
	if err := keyboard.Update(); err != nil {
		return err
	}

	pr := g.parent.Rotation()
	pr += 1.0
	if pr >= 360.0 {
		pr = 0.0
	}
	g.parent.SetRotation(pr)

	return nil
}
