package ecs

import "github.com/google/uuid"

type entity struct {
	id         uuid.UUID
	components map[ECSTypeID]IComponent
	childMap   map[uuid.UUID]IEntity
	children   []IEntity
	parent     IEntity
}

// NewEntity creates a new entity with a unique ID.
//
// Disposing of an entity will also dispose of all its components and children.
func NewEntity() IEntity {
	e := &entity{
		id:         uuid.New(),
		components: make(map[ECSTypeID]IComponent),
		childMap:   make(map[uuid.UUID]IEntity),
		children:   []IEntity{},
		parent:     nil,
	}
	return e
}

func (e *entity) Dispose() {
	for _, component := range e.components {
		component.Dispose()
	}
	for _, child := range e.children {
		child.Dispose()
	}

	e.components = nil
	e.children = nil
	e.childMap = nil
	e.parent = nil

	e.id = uuid.Nil
}

func (e *entity) ID() uuid.UUID {
	return e.id
}

func (e *entity) AddComponent(component IComponent) IEntity {
	if component == nil {
		panic("Cannot add nil component to entity")
	}
	if e.HasComponent(component.GetTypeID()) {
		panic("Component already exists in entity")
	}
	e.components[component.GetTypeID()] = component
	return e
}

func (e *entity) AddComponents(components ...IComponent) IEntity {
	for _, component := range components {
		e.AddComponent(component)
	}
	return e
}

func (e *entity) RemoveComponent(typeID ECSTypeID) IEntity {
	if _, exists := e.components[typeID]; !exists {
		return e
	}
	delete(e.components, typeID)
	return e
}

func (e *entity) RemoveComponents(typeIDs ...ECSTypeID) IEntity {
	if len(typeIDs) == 0 {
		return e
	}

	for _, typeID := range typeIDs {
		delete(e.components, typeID)
	}

	return e
}

func (e *entity) GetComponent(typeID ECSTypeID) IComponent {
	component, exists := e.components[typeID]
	if !exists {
		return nil
	}
	return component
}

func (e *entity) GetComponents(typeIDs ...ECSTypeID) []IComponent {
	if len(typeIDs) == 0 {
		return nil
	}

	components := make([]IComponent, 0, len(typeIDs))
	for _, typeID := range typeIDs {
		if component, exists := e.components[typeID]; exists {
			components = append(components, component)
		}
	}

	return components
}

func (e *entity) HasComponent(typeID ECSTypeID) bool {
	_, exists := e.components[typeID]
	return exists
}

func (e *entity) HasComponents(typeIDs ...ECSTypeID) bool {
	if len(typeIDs) == 0 {
		return true
	}

	for _, typeID := range typeIDs {
		if !e.HasComponent(typeID) {
			return false
		}
	}

	return true
}

func (e *entity) AddChild(child IEntity) bool {
	if child == nil {
		panic("Cannot add nil child to entity")
	}

	if e.HasChild(child) {
		return false
	}

	if child.GetParent() != nil {
		child.GetParent().RemoveChild(child)
	}

	child.internal_SetParent(e)

	e.children = append(e.children, child)
	e.childMap[child.ID()] = child

	return true
}

func (e *entity) RemoveChild(child IEntity) bool {
	if child == nil || !e.HasChild(child) {
		return false
	}

	child.internal_SetParent(nil)

	for i, c := range e.children {
		if c.ID() == child.ID() {
			e.children = append(e.children[:i], e.children[i+1:]...)
			break
		}
	}
	delete(e.childMap, child.ID())

	return true
}

func (e *entity) Children() []IEntity {
	return e.children
}

func (e *entity) GetParent() IEntity {
	return e.parent
}

func (e *entity) GetChildByID(id uuid.UUID) IEntity {
	if child, exists := e.childMap[id]; exists {
		return child
	}
	return nil
}

func (e *entity) HasChild(child IEntity) bool {
	if child == nil {
		return false
	}
	if parent := child.GetParent(); parent != nil && parent.ID() == e.ID() {
		return true
	}
	return false
}

func (e *entity) HasChildByID(id uuid.UUID) bool {
	return e.HasChild(e.GetChildByID(id))
}

func (e *entity) internal_SetParent(parent IEntity) {
	for _, component := range e.components {
		component.OnParentChanged(e.parent, parent)
	}
	e.parent = parent
}
