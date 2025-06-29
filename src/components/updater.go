package components

type Updater interface {
	Update() error
}

type updater struct {
}

func NewUpdater() Updater {
	return &updater{}
}

func (u *updater) Update() error {
	return nil
}
