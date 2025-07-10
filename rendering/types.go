package rendering

import "github.com/hajimehoshi/ebiten/v2"

type IAnchorPoint interface {
	GetAnchor() (float64, float64) // Retrieves the anchor point of a renderer
	SetAnchor(x, y float64)        // Sets the anchor point of a renderer
}

type IRenderer interface {
	Render(screen *ebiten.Image, camera ebiten.GeoM, matrix ebiten.GeoM) // Renders the component to the screen using the provided camera transformation
	Dispose()                                                            // Disposes of the renderer, releasing any resources it holds
}

type ISpriteRenderer interface {
	IAnchorPoint
	IRenderer

	SetImage(image *ebiten.Image) // Sets the image to be rendered by the sprite renderer
	GetImage() *ebiten.Image      // Retrieves the current image being rendered

	SetBlendMode(mode ebiten.Blend)     // SetBlendMode sets the blend mode for rendering
	SetColorScale(cs ebiten.ColorScale) // SetColorScale sets the color scale for rendering
	SetFilter(filter ebiten.Filter)     // SetFilter sets the filter for rendering
	SetMipMaps(toggle bool)             // UseMipMaps enables or disables mipmaps for the renderer
}
