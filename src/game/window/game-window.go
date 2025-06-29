package window

type Window interface {
	GetScreenSize() (width, height int)

	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type gamewindow struct {
	width, height int
}

func NewGameWindow(screenWidth, screenHeight int) Window {
	return &gamewindow{
		width:  screenWidth,
		height: screenHeight,
	}
}

func (g *gamewindow) GetScreenSize() (width, height int) {
	return g.width, g.height
}

func (g *gamewindow) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = g.width, g.height

	if outsideWidth < g.width {
		screenWidth = outsideWidth
	}

	if outsideHeight < g.height {
		screenHeight = outsideHeight
	}

	return screenWidth, screenHeight
}
