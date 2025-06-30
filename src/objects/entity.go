package objects

import "github.com/ADM87/ggame/src/components"

// Entity defines a basic interface for an object within the game world
type Entity interface {
	Transform() components.Transform // Transform returns the Transform interface for position and transformation capabilities
}

type entity struct {
	transform components.Transform
}

func NewEntity() Entity {
	return &entity{
		transform: components.NewTransform(0, 0),
	}
}

func (e *entity) Transform() components.Transform {
	return e.transform
}
