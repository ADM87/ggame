package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var op = ebiten.DrawImageOptions{
	Filter: ebiten.FilterNearest,
}

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	view := g.gameCamera.GetViewMatrix()

	var mat ebiten.GeoM
	for _, actor := range g.actors {
		mat = actor.Transform().Matrix()
		drawImage(renderTarget, actor.Renderer().GetImage(), &view, &mat)
	}

	mat = g.player.Transform().Matrix()
	drawImage(renderTarget, g.player.Renderer().GetImage(), &view, &mat)

	ebitenutil.DebugPrint(renderTarget, "Press ESC to exit")
}

func drawImage(renderTarget *ebiten.Image, image *ebiten.Image, view *ebiten.GeoM, matrix *ebiten.GeoM) {
	op.GeoM.Reset()
	op.GeoM.Concat(*view)
	op.GeoM.Concat(*matrix)
	renderTarget.DrawImage(image, &op)
}
