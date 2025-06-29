package game

import (
	"github.com/ADM87/ggame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	view := g.gameCamera.GetViewMatrix()

	img, err := resources.LoadImage("tile_0000")
	if err != nil {
		ebitenutil.DebugPrint(renderTarget, "Error loading image: "+err.Error())
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM = view

	renderTarget.DrawImage(img, op)

	ebitenutil.DebugPrint(renderTarget, "Press ESC to exit")
}
