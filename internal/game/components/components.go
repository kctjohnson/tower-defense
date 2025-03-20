package components

import "ecstemplate/pkg/ecs"

const (
	Position   ecs.ComponentType = "position"
	Health     ecs.ComponentType = "health"
	Velocity   ecs.ComponentType = "velocity"
	Enemy      ecs.ComponentType = "enemy"
	Tower      ecs.ComponentType = "tower"
	Projectile ecs.ComponentType = "projectile"
	Path       ecs.ComponentType = "path"
	Wallet     ecs.ComponentType = "wallet"
	Renderable ecs.ComponentType = "renderable"
)

type PositionComponent struct {
	ecs.Component
	X, Y float64
}

func (c PositionComponent) GetType() ecs.ComponentType {
	return Position
}

type HealthComponent struct {
	ecs.Component
	Current, Max float64
}

func (c HealthComponent) GetType() ecs.ComponentType {
	return Health
}

type VelocityComponent struct {
	ecs.Component
	X, Y float64
}

func (c VelocityComponent) GetType() ecs.ComponentType {
	return Velocity
}

type EnemyComponent struct {
	ecs.Component
	Type   string
	Speed  float64
	Reward float64
}

func (c EnemyComponent) GetType() ecs.ComponentType {
	return Enemy
}

type TowerComponent struct {
	ecs.Component
	Damage, Range, Cooldown, LastFired float64
}

func (c TowerComponent) GetType() ecs.ComponentType {
	return Tower
}

type ProjectileComponent struct {
	ecs.Component
	Damage, Speed, Reward float64
}

func (c ProjectileComponent) GetType() ecs.ComponentType {
	return Projectile
}

type PathComponent struct {
	ecs.Component
	Waypoints []PositionComponent
}

func (c PathComponent) GetType() ecs.ComponentType {
	return Path
}

type WalletComponent struct {
	ecs.Component
	Money float64
}

func (c WalletComponent) GetType() ecs.ComponentType {
	return Wallet
}

type RenderableComponent struct {
	ecs.Component
	Symbol string
}

func (c RenderableComponent) GetType() ecs.ComponentType {
	return Renderable
}

var ComponentTypes = []ecs.ComponentType{
	Position,
	Health,
	Velocity,
	Enemy,
	Tower,
	Projectile,
	Path,
	Wallet,
	Renderable,
}
