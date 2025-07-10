package ecs

type component struct {
	IIdenifyableType

	owner IEntity
}

func NewComponent(typeID ECSTypeID) IComponent {
	return NewComponentWithOwner(typeID, NewEntity())
}

func NewComponentWithOwner(typeID ECSTypeID, owner IEntity) IComponent {
	if typeID == 0 {
		panic("ComponentTypeID cannot be zero")
	}
	if owner == nil {
		panic("Component owner cannot be nil")
	}
	return &component{
		IIdenifyableType: NewIdentifyableType(typeID),
		owner:            owner,
	}
}

func (c *component) GetEntity() IEntity {
	return c.owner
}

func (c *component) OnParentChanged(oldParent IEntity, newParent IEntity) {
}

func (c *component) Dispose() {
	c.owner = nil
}
