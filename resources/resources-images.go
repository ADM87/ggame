package resources

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ADM87/ggame/sys"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DefaultImage returns the default image resource.
//
// Panics if the default image cannot be loaded.
func DefaultImage() *ebiten.Image {
	return LoadImage(DefaultImageResource)
}

// LoadImage loads an image resource by its name from the static or dynamic resources.
//
// LoadImage will first check if the requested image has already been loaded and cached.
func LoadImage(name string) *ebiten.Image {
	if img, exists := imageCache[name]; exists {
		return img
	}

	var file fs.File
	var err error

	if resource, exists := manifest.Static[name]; exists {
		path := filepath.Join("static", resource.Path)
		file, err = staticResources.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				sys.Logger().Warnf("Static resource '%s' not found, falling back to default image resource", name)
				return DefaultImage()
			}
			panic("Failed to open static resource: " + err.Error())
		}
	} else if resource, exists := manifest.Dynamic[name]; exists {
		path := filepath.Join(resourceFolder, "dynamic", resource.Path)
		file, err = os.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				sys.Logger().Warnf("Dynamic resource '%s' not found, falling back to default image resource", name)
				return DefaultImage()
			}
			panic("Failed to open dynamic resource: " + err.Error())
		}
	} else {
		sys.Logger().Warnf("Resource '%s' not found in manifest, falling back to default image resource", name)
		return DefaultImage()
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic("Failed to close file: " + err.Error())
		}
	}()

	img, _, err := ebitenutil.NewImageFromReader(file)
	if err != nil {
		panic("Failed to load image from reader: " + err.Error())
	}
	imageCache[name] = img

	sys.Logger().Debugf("Loaded image resource '%s'", name)
	return img
}

func UnloadImage(name string) error {
	if _, exists := imageCache[name]; !exists {
		sys.Logger().Warnf("Cannot unload resource '%s', it is not loaded", name)
		return nil
	}

	if _, exists := manifest.Static[name]; exists {
		sys.Logger().Warnf("Cannot unload static resource '%s', it is a static resource", name)
		return nil
	}

	delete(imageCache, name)
	return nil
}
