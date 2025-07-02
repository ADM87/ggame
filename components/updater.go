package components

type updater[T IUpdatable] struct {
	updateFuncs []UpdateFunc[T] // List of update functions for components
}

func NewUpdater[T IUpdatable]() IUpdater[T] {
	return &updater[T]{
		updateFuncs: make([]UpdateFunc[T], 0),
	}
}

func (u *updater[T]) Add(updateFunc UpdateFunc[T]) {
	u.updateFuncs = append(u.updateFuncs, updateFunc)
}

func (u *updater[T]) Remove(updateFunc UpdateFunc[T]) {
}

func (u *updater[T]) Update(dt float64) {
	for _, updateFunc := range u.updateFuncs {
		updateFunc(dt)
	}
}
