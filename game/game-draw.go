package game

import (
	"fmt"

	"github.com/ADM87/ggame/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	renderables = make([]components.IRenderer, 0)
)

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	for _, r := range renderables {
		r.Render(renderTarget, ebiten.GeoM{}, ebiten.GeoM{})
	}

	ebitenutil.DebugPrint(renderTarget, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}
