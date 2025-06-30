package components

type ComponentID int

const (
	RenderComponentID    ComponentID = iota + 1 // RenderComponentID is the ID for the render component
	SpriteRendererID                            // SpriteRendererID is the ID for the sprite renderer component
	TransformComponentID                        // TransformComponentID is the ID for the transform component
)

func (id ComponentID) String() string {
	switch id {
	case RenderComponentID:
		return "RenderComponentID"
	case SpriteRendererID:
		return "SpriteRendererID"
	case TransformComponentID:
		return "TransformComponentID"
	default:
		return "UnknownComponentID"
	}
}

type Component interface {
	ID() ComponentID // ID returns the unique identifier for the component
}

type component struct {
	id ComponentID // id is the unique identifier for the component
}

// NewComponent creates a new component with the specified ID
func NewComponent(id ComponentID) Component {
	return &component{
		id: id,
	}
}

func (c *component) ID() ComponentID {
	return c.id // ID returns the unique identifier for the component
}
