package game

import (
	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/ecs"
)

func NewActor() components.IActor {
	actor := ecs.NewEntity()
	actor.AddComponents(
		components.NewSpriteRenderComponentWithOwner(actor),
		components.NewTransformComponentWithOwner(actor),
		components.NewActorComponentWithOwner(actor))
	return actor.GetComponent(components.ActorTypeID).(components.IActor)
}

func NewCamera() components.ICamera {
	camera := ecs.NewEntity()
	camera.AddComponents(
		components.NewCameraComponentWithOwner(camera, int(ScreenWidth), int(ScreenHeight)),
		components.NewTransformComponentWithOwner(camera))
	return camera.GetComponent(components.CameraTypeID).(components.ICamera)
}
