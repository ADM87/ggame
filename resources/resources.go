package resources

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed static/*
var StaticResources embed.FS

var imageCache = make(map[string]*ebiten.Image)

var manifest *ResourceManifest

var resourceFolder = "resources"

func LoadImage(name string) (*ebiten.Image, error) {
	if img, exists := imageCache[name]; exists {
		return img, nil
	}

	if manifest == nil {
		return nil, os.ErrNotExist
	}

	var file fs.File
	var err error

	if resource, exists := manifest.Static[name]; exists {
		path := filepath.Join("static", resource.Path)
		file, err = StaticResources.Open(path)
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
		return nil, os.ErrNotExist
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

func UnloadImage(path string) error {
	if img, exists := imageCache[path]; exists {
		delete(imageCache, path)
		img.Deallocate()
	}
	return nil
}

func Initialize(rootDir string) error {
	resPath, err := filepath.Abs(filepath.Join(rootDir, resourceFolder))
	if err != nil {
		return err
	}
	resourceFolder = resPath

	manifestFile, err := StaticResources.Open("static/manifest.json")
	if err != nil {
		return err
	}
	defer manifestFile.Close()

	manifestData, err := StaticResources.ReadFile("static/manifest.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return err
	}

	for name := range manifest.Static {
		img, err := LoadImage(name)
		if err != nil {
			return err
		}
		imageCache[name] = img
	}

	return nil
}
