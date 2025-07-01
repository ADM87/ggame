package game

import "github.com/ADM87/ggame/keyboard"

func (g *gameshell) Update() error {
	if err := keyboard.Update(); err != nil {
		return err
	}
	return nil
}
