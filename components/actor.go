package components

import "github.com/ADM87/ggame/ecs"

type actorComponent struct {
	ecs.IComponent

	transform       ITransformComponent
	renderComponent IRenderComponent
}

func NewActorComponent() IActor {
	return NewActorComponentWithOwner(ecs.NewEntity())
}

func NewActorComponentWithOwner(owner ecs.IEntity) IActor {
	return &actorComponent{
		IComponent:      ecs.NewComponentWithOwner(ActorTypeID, owner),
		transform:       nil,
		renderComponent: nil,
	}
}

func (p *actorComponent) Transform() ITransformComponent {
	return p.GetEntity().GetComponent(TransformTypeID).(ITransformComponent)
}

func (p *actorComponent) RenderComponent() IRenderComponent {
	return p.GetEntity().GetComponent(RenderTypeID).(IRenderComponent)
}
