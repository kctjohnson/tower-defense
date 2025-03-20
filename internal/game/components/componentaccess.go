package components

import "ecstemplate/pkg/ecs"

type ComponentAccess struct {
	world *ecs.World
}

func NewComponentAccess(world *ecs.World) *ComponentAccess {
	return &ComponentAccess{world: world}
}

func (c *ComponentAccess) GetSampleComponent(entity ecs.Entity) (*SampleComponent, bool) {
	component, found := c.world.ComponentManager.GetComponent(entity, Sample)
	if !found {
		return nil, false
	}
	return component.(*SampleComponent), true
}
