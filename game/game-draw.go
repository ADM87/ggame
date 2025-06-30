package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var op = &ebiten.DrawImageOptions{
	Filter: ebiten.FilterNearest,
}
var identityMatrix = ebiten.GeoM{}

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	viewMatrix := g.gameCamera.GetViewMatrix()
	for _, entity := range g.entities {
		entity.Render(renderTarget, viewMatrix, identityMatrix, op)
	}

	ebitenutil.DebugPrint(renderTarget, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}
