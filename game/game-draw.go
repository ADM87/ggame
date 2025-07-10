package game

import (
	"fmt"
	"strings"

	"github.com/ADM87/ggame/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var sb = strings.Builder{}

func (g *gameshell) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(backgroundColor)

	if renderSystem, ok := g.systems[systems.RenderSystemTypeID].(systems.IRenderSystem); ok {
		renderSystem.Render(renderTarget, g.camera.ViewMatrix())
	}

	ebitenutil.DebugPrint(renderTarget, g.GetDebugInfo())
}

func (g *gameshell) GetDebugInfo() string {
	sb.Reset()

	camX, camY := g.camera.GetPosition()
	sb.WriteString(fmt.Sprintf("Camera Position: (%.2f, %.2f)\n", camX, camY))

	zoom := g.camera.GetZoom()
	sb.WriteString(fmt.Sprintf("Camera Zoom: %.2f\n", zoom))

	return sb.String()
}
