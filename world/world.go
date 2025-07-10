package world

import (
	"github.com/ADM87/ggame/ecs"
	"github.com/google/uuid"
)

var entities = make(map[uuid.UUID]ecs.IEntity)

var entityComponents = make(map[uuid.UUID]map[ecs.ECSTypeID]ecs.IComponent)

func AddEntity(entity ecs.IEntity) {
	if entity == nil {
		panic("Cannot add nil entity")
	}
	if _, exists := entities[entity.ID()]; exists {
		panic("Entity with ID " + entity.ID().String() + " already exists")
	}
	if entity.ID() == uuid.Nil {
		panic("Entity ID cannot be nil")
	}
	if _, exists := entityComponents[entity.ID()]; exists {
		panic("Entity with ID " + entity.ID().String() + " already has components registered with the world")
	}

	entityComponents[entity.ID()] = make(map[ecs.ECSTypeID]ecs.IComponent)
	for _, component := range entity.GetComponents() {
		entityComponents[entity.ID()][component.GetTypeID()] = component
	}

	entities[entity.ID()] = entity
}

func RemoveEntity(entity ecs.IEntity) {
	if entity == nil {
		panic("Cannot remove nil entity")
	}
	if _, exists := entities[entity.ID()]; !exists {
		panic("Entity with ID " + entity.ID().String() + " does not exist")
	}
	if entity.ID() == uuid.Nil {
		panic("Cannot remove entity with nil ID")
	}

	for _, component := range entity.GetComponents() {
		entityComponents[entity.ID()][component.GetTypeID()] = nil
	}

	delete(entityComponents, entity.ID())
	delete(entities, entity.ID())
}

func RemoveEntityByID(id uuid.UUID) {
	if id == uuid.Nil {
		panic("Cannot remove entity with nil ID")
	}

	entity, exists := entities[id]
	if !exists {
		panic("Entity with ID " + id.String() + " does not exist")
	}

	entity.Dispose()
	delete(entities, id)
}

func GetEntity(id uuid.UUID) (ecs.IEntity, bool) {
	entity, exists := entities[id]
	return entity, exists
}

func FilterEntities(filter ...ecs.ECSTypeID) []ecs.IEntity {
	if len(filter) == 0 {
		return nil
	}
	filtered := make([]ecs.IEntity, 0, len(entities))
	for _, entity := range entities {
		if entity.HasComponents(filter...) {
			filtered = append(filtered, entity)
		}
	}
	return filtered
}
