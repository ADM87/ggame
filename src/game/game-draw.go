package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var op = &ebiten.DrawImageOptions{
	Filter: ebiten.FilterNearest,
}

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	view := g.gameCamera.GetViewMatrix()
	for _, actor := range g.actors {
		actor.Renderer().Render(renderTarget, view, actor.Transform().Matrix(), op)
	}

	ebitenutil.DebugPrint(renderTarget, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}
