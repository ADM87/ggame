package components

import (
	"math"

	"github.com/ADM87/ggame/ecs"
	"github.com/ADM87/ggame/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

type transformComponent struct {
	ecs.IComponent

	x, y     float64 // Position in local space
	ox, oy   float64 // Origin in local space (pivot point)
	sx, sy   float64 // Scale factors in local space
	rot, deg float64 // Rotation in radians

	localMatrix ebiten.GeoM
	localDirty  bool

	worldMatrix ebiten.GeoM
	worldDirty  bool

	parentTransform ITransformComponent // Reference to the parent transform component, if any
}

func NewTransformComponent() ITransformComponent {
	return NewTransformComponentWithOwner(ecs.NewEntity())
}

func NewTransformComponentWithOwner(owner ecs.IEntity) ITransformComponent {
	return &transformComponent{
		IComponent:      ecs.NewComponentWithOwner(TransformTypeID, owner),
		x:               0,
		y:               0,
		ox:              0,
		oy:              0,
		sx:              1,
		sy:              1,
		rot:             0,
		deg:             0,
		localMatrix:     ebiten.GeoM{},
		localDirty:      true,
		worldMatrix:     ebiten.GeoM{},
		worldDirty:      true,
		parentTransform: nil,
	}
}

func (t *transformComponent) OnParentChanged(oldParent ecs.IEntity, newParent ecs.IEntity) {
	t.internal_CacheParentTransform()
}

func (t *transformComponent) GetPosition() (x, y float64) {
	return t.x, t.y
}

func (t *transformComponent) GetRotation() float64 {
	return t.deg
}

func (t *transformComponent) GetScale() (x, y float64) {
	return t.sx, t.sy
}

func (t *transformComponent) SetPosition(x, y float64) {
	if t.x != x || t.y != y {
		t.x = x
		t.y = y
		t.internal_SetDirty(true, true)
	}
}

func (t *transformComponent) SetRotation(degrees float64) {
	wrapped := math.Mod(degrees, 360.0)
	if wrapped < 0 {
		wrapped += 360.0
	}
	if t.deg != wrapped {
		t.deg = wrapped
		t.rot = wrapped * sys.DegToRads
		t.internal_SetDirty(true, true)
	}
}

func (t *transformComponent) SetScale(x, y float64) {
	if t.sx != x || t.sy != y {
		t.sx = x
		t.sy = y
		t.internal_SetDirty(true, true)
	}
}

func (t *transformComponent) LocalMatrix() ebiten.GeoM {
	if t.localDirty {
		t.localMatrix.Reset()
		t.localMatrix.Translate(-t.ox, -t.oy) // Translate to origin first
		t.localMatrix.Scale(t.sx, t.sy)       // Apply scaling first
		t.localMatrix.Rotate(t.rot)           // Apply rotation second
		t.localMatrix.Translate(t.x, t.y)     // Apply translation last
		t.localDirty = false
	}
	return t.localMatrix
}

func (t *transformComponent) WorldMatrix() ebiten.GeoM {
	if t.worldDirty {
		localMatrix := t.LocalMatrix()
		if t.parentTransform != nil {
			t.worldMatrix = t.parentTransform.WorldMatrix()
			t.worldMatrix.Concat(localMatrix)
		} else {
			t.worldMatrix = localMatrix // No parent, use local matrix directly
		}
		t.worldDirty = false
	}
	return t.worldMatrix
}

func (t *transformComponent) internal_SetDirty(local, world bool) {
	t.localDirty = t.localDirty || local
	t.worldDirty = t.worldDirty || world
	if world {
		for _, child := range t.GetEntity().Children() {
			if childTransform, ok := child.GetComponent(TransformTypeID).(ITransformComponent); ok {
				childTransform.internal_SetDirty(false, true)
			}
		}
	}
}

func (t *transformComponent) internal_CacheParentTransform() {
	t.parentTransform = nil
	if parent := t.GetEntity().GetParent(); parent != nil {
		if parentTransform, ok := parent.GetComponent(TransformTypeID).(ITransformComponent); ok {
			t.parentTransform = parentTransform
		}
	}
	t.internal_SetDirty(false, true)
}
