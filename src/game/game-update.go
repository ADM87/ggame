package game

import "github.com/ADM87/ggame/src/keyboard"

func (g *gameshell) Update() error {
	if err := keyboard.Update(); err != nil {
		return err
	}
	g.gameCamera.MoveTo(g.player.Transform().Position())
	return nil
}
