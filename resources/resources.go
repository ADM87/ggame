package resources

import (
	"embed"
	"encoding/json"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed static/*
	staticResources embed.FS
	manifest        ResourceManifest

	imageCache       = make(map[string]*ebiten.Image)
	resourceFolder   = "resources"
	missingImageName = "10x10"
)

func Initialize(rootDir string) error {
	resPath, err := filepath.Abs(filepath.Join(rootDir, resourceFolder))
	if err != nil {
		return err
	}
	resourceFolder = resPath

	manifestData, err := staticResources.ReadFile("static/manifest.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return err
	}

	for name := range manifest.Static {
		_, err := LoadImage(name)
		if err != nil {
			return err
		}
	}

	return nil
}
