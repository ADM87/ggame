package resources

import (
	"embed"
	"encoding/json"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DefaultImageResource = "10x10"
)

var (
	//go:embed static/*
	staticResources embed.FS
	manifest        ResourceManifest

	imageCache     = make(map[string]*ebiten.Image)
	resourceFolder = "resources"
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

	// Load static images so they are ready for use from the start
	for name := range manifest.Static {
		LoadImage(name)
	}

	return nil
}
