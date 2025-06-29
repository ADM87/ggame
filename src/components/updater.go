package components

// Updater defines the interface for components that can update their state
type Updater interface {
	Update() error // Update performs the update logic for the component
}

type updater struct {
}

// NewUpdater creates a basic updater component
func NewUpdater() Updater {
	return &updater{}
}

func (u *updater) Update() error {
	return nil
}
