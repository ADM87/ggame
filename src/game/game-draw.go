package game

import (
	"github.com/ADM87/ggame/src/components"
	"github.com/hajimehoshi/ebiten/v2"
)

var op = ebiten.DrawImageOptions{
	Filter: ebiten.FilterNearest,
}

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	view := g.gameCamera.GetViewMatrix()
	for _, actor := range g.actors {
		drawImage(renderTarget, actor.Renderer().GetImage(), view, actor.Transform())
	}
}

func drawImage(renderTarget *ebiten.Image, image *ebiten.Image, view ebiten.GeoM, transform components.Transform) {
	op.GeoM.Reset()
	op.GeoM.Concat(transform.Matrix())
	op.GeoM.Concat(view)
	renderTarget.DrawImage(image, &op)
}
