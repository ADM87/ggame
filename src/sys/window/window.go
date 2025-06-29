package window

type Window interface {
	GetWidth() int
	GetHeight() int
	Layout(windowWidth, windowHeight int) (screenWidth, screenHeight int)
}

type syswindow struct {
	width  int
	height int
}

func NewWindow(width, height int) Window {
	return &syswindow{
		width:  width,
		height: height,
	}
}

func (w *syswindow) GetWidth() int {
	return w.width
}

func (w *syswindow) GetHeight() int {
	return w.height
}

func (w *syswindow) Layout(windowWidth, windowHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = w.width, w.height

	if windowWidth < w.width {
		screenWidth = windowWidth
	}

	if windowHeight < w.height {
		screenHeight = windowHeight
	}

	return screenWidth, screenHeight
}
