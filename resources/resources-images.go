package resources

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ADM87/ggame/sys"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// LoadImage loads an image resource by its name from the static or dynamic resources.
//
// LoadImage will first check if the requested image has already been loaded and cached.
func LoadImage(name string) (*ebiten.Image, error) {
	if img, exists := imageCache[name]; exists {
		return img, nil
	}

	var file fs.File
	var err error

	if resource, exists := manifest.Static[name]; exists {
		path := filepath.Join("static", resource.Path)
		file, err = staticResources.Open(path)
		if err != nil {
			return nil, err
		}
	} else if resource, exists := manifest.Dynamic[name]; exists {
		path := filepath.Join(resourceFolder, "dynamic", resource.Path)
		file, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	} else {
		sys.Logger().Warnf("Resource '%s' not found in manifest, falling back to missing image resource", name)
		return LoadImage(missingImageName)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic("Failed to close file: " + err.Error())
		}
	}()

	img, _, err := ebitenutil.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	imageCache[name] = img

	return img, nil
}

func UnloadImage(name string) error {
	if _, exists := imageCache[name]; !exists {
		sys.Logger().Warnf("Cannot unload resource '%s', it is not loaded", name)
		return nil
	}

	if _, exists := manifest.Static[name]; exists {
		sys.Logger().Warnf("Cannot unload static resource '%s', it is part of the static manifest", name)
		return nil
	}

	delete(imageCache, name)
	return nil
}
