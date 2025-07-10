package ecs

import (
	"hash/fnv"
	"reflect"

	"github.com/google/uuid"
)

func GetTypeID[T any]() ECSTypeID {
	typeName := reflect.TypeOf((*T)(nil)).Elem().String()
	if typeName == "" {
		panic("Component type must have a name")
	}
	h := fnv.New64a()
	h.Write([]byte(typeName))
	return ECSTypeID(h.Sum64())
}

func NewEntityID() uuid.UUID {
	return uuid.New()
}
