package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	for _, entity := range g.entities {
		entity.Render(renderTarget, ebiten.GeoM{}, ebiten.GeoM{})
	}

	ebitenutil.DebugPrint(renderTarget, fmt.Sprintf("%0.2f", g.parent.Rotation()))
}
