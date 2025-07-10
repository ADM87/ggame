package ecs

type system struct {
	IIdenifyableType

	filters []ECSTypeID
}

func NewSystem(typeID ECSTypeID, filters ...ECSTypeID) ISystem {
	if typeID == 0 {
		panic("SystemTypeID cannot be zero")
	}
	if len(filters) == 0 {
		panic("System must have at least one filter")
	}
	return &system{
		IIdenifyableType: NewIdentifyableType(typeID),
		filters:          filters,
	}
}

func (s *system) Filter() []ECSTypeID {
	if s.filters == nil {
		return nil
	}
	filters := make([]ECSTypeID, len(s.filters))
	copy(filters, s.filters)
	return filters
}

func (s *system) Update(entities []IEntity, dt float64) {

}

func (s *system) Dispose() {
	s.filters = nil
}
