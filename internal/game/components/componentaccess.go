package components

import "ecstemplate/pkg/ecs"

type ComponentAccess struct {
	world *ecs.World
}

func NewComponentAccess(world *ecs.World) *ComponentAccess {
	return &ComponentAccess{world: world}
}

func GetComponentT[T ecs.ComponentInterface](
	world *ecs.World,
	entity ecs.Entity,
	componentType ecs.ComponentType,
) (T, bool) {
	component, found := world.ComponentManager.GetComponent(entity, componentType)
	if !found {
		var zero T
		return zero, false
	}

	return component.(T), true
}

func (c *ComponentAccess) GetPositionComponent(entity ecs.Entity) (*PositionComponent, bool) {
	return GetComponentT[*PositionComponent](c.world, entity, Position)
}

func (c *ComponentAccess) GetHealthComponent(entity ecs.Entity) (*HealthComponent, bool) {
	return GetComponentT[*HealthComponent](c.world, entity, Health)
}

func (c *ComponentAccess) GetVelocityComponent(entity ecs.Entity) (*VelocityComponent, bool) {
	return GetComponentT[*VelocityComponent](c.world, entity, Velocity)
}

func (c *ComponentAccess) GetEnemyComponent(entity ecs.Entity) (*EnemyComponent, bool) {
	return GetComponentT[*EnemyComponent](c.world, entity, Enemy)
}

func (c *ComponentAccess) GetTowerComponent(entity ecs.Entity) (*TowerComponent, bool) {
	return GetComponentT[*TowerComponent](c.world, entity, Tower)
}

func (c *ComponentAccess) GetProjectileComponent(entity ecs.Entity) (*ProjectileComponent, bool) {
	return GetComponentT[*ProjectileComponent](c.world, entity, Projectile)
}

func (c *ComponentAccess) GetPathComponent(entity ecs.Entity) (*PathComponent, bool) {
	return GetComponentT[*PathComponent](c.world, entity, Path)
}

func (c *ComponentAccess) GetPathFollowComponent(entity ecs.Entity) (*PathFollowComponent, bool) {
	return GetComponentT[*PathFollowComponent](c.world, entity, PathFollow)
}

func (c *ComponentAccess) GetWalletComponent(entity ecs.Entity) (*WalletComponent, bool) {
	return GetComponentT[*WalletComponent](c.world, entity, Wallet)
}

func (c *ComponentAccess) GetRenderableComponent(entity ecs.Entity) (*RenderableComponent, bool) {
	return GetComponentT[*RenderableComponent](c.world, entity, Renderable)
}
