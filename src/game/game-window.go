package game

type Window interface {
	GetWidth() int
	GetHeight() int

	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type gamewindow struct {
	width  int
	height int
}

func NewGameWindow(width, height int) Window {
	return &gamewindow{
		width:  width,
		height: height,
	}
}

func (g *gamewindow) GetWidth() int {
	return g.width
}

func (g *gamewindow) GetHeight() int {
	return g.height
}

func (g *gamewindow) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
