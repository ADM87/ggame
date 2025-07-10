package ecs

type identifyableType struct {
	typeID ECSTypeID
}

func NewIdentifyableType(typeID ECSTypeID) IIdenifyableType {
	if typeID == 0 {
		panic("TypeID cannot be zero")
	}
	return &identifyableType{typeID: typeID}
}

func (t *identifyableType) GetTypeID() ECSTypeID {
	return t.typeID
}
