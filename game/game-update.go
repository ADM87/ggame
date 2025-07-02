package game

import (
	"time"

	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/entities"
	"github.com/ADM87/ggame/keyboard"
)

var (
	entityUpdater = components.NewUpdater[entities.IEntity]()
	maxDeltaTime  = 1.0 / 30.0 // Cap delta time to 30 FPS minimum (33ms max)
	lastTime      = time.Time{}
	targetDT      = 1.0 / 60.0 // Target 60 FPS (16.67ms)
)

func (g *gameshell) Update() error {
	if err := keyboard.Update(); err != nil {
		return err
	}

	currentTime := time.Now()

	// Initialize lastTime on first frame
	if lastTime.IsZero() {
		lastTime = currentTime
		return nil // Skip first frame to establish baseline
	}

	// Calculate actual delta time using high-resolution timer
	dt := currentTime.Sub(lastTime).Seconds()
	lastTime = currentTime

	// Cap delta time to prevent large jumps during startup/frame drops
	if dt > maxDeltaTime {
		dt = maxDeltaTime
	}

	// Use target delta time if calculated dt is too small (unrealistic high FPS)
	if dt < targetDT/2 {
		dt = targetDT
	}

	entityUpdater.Update(dt)

	return nil
}
