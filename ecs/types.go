package ecs

import (
	"github.com/google/uuid"
)

type ECSTypeID uint64

type IIdenifyableType interface {
	GetTypeID() ECSTypeID // Retrieves the ECSTypeID
}

type IEntity interface {
	ID() uuid.UUID                     // Unique identifier for the entity
	GetChildByID(id uuid.UUID) IEntity // Retrieves a child entity by its ID
	HasChildByID(id uuid.UUID) bool    // Checks if the entity has a child with the specified ID

	AddComponent(component IComponent) IEntity      // Adds a component to the entity
	AddComponents(components ...IComponent) IEntity // Adds multiple components to the entity

	RemoveComponent(typeID ECSTypeID) IEntity      // Removes a component by its type ID
	RemoveComponents(typeIDs ...ECSTypeID) IEntity // Removes multiple components by their type IDs

	GetComponent(typeID ECSTypeID) IComponent        // Retrieves a component by its type ID
	GetComponents(typeIDs ...ECSTypeID) []IComponent // Retrieves multiple components by their type IDs

	HasComponent(typeID ECSTypeID) bool      // Checks if the entity has a component of the specified type
	HasComponents(typeIDs ...ECSTypeID) bool // Checks if the entity has all specified components

	AddChild(child IEntity) bool    // Adds a child entity to this entity
	RemoveChild(child IEntity) bool // Removes a child entity from this entity
	HasChild(child IEntity) bool    // Checks if the entity has a specific child entity
	Children() []IEntity            // Retrieves all child entities of this entity
	GetParent() IEntity             // Retrieves the parent entity of this entity

	Dispose()

	internal_SetParent(parent IEntity)
}

type IComponent interface {
	IIdenifyableType

	GetEntity() IEntity // Retrieves the owner entity of this component

	OnParentChanged(oldParent IEntity, newParent IEntity) // Called when the parent of the component changes

	Dispose()
}

type ISystem interface {
	IIdenifyableType

	Filter() []ECSTypeID
	Update(entities []IEntity, dt float64)

	Dispose()
}
