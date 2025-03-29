package components

import (
	"time"

	"ecstemplate/pkg/ecs"
)

const (
	Display           ecs.ComponentType = "display"
	Cursor            ecs.ComponentType = "cursor"
	GameState         ecs.ComponentType = "game_state"
	Player            ecs.ComponentType = "player"
	Enemy             ecs.ComponentType = "enemy"
	BoundingBox       ecs.ComponentType = "bounding_box"
	Position          ecs.ComponentType = "position"
	Health            ecs.ComponentType = "health"
	Velocity          ecs.ComponentType = "velocity"
	Tower             ecs.ComponentType = "tower"
	TowerTemplate     ecs.ComponentType = "tower_template"
	Projectile        ecs.ComponentType = "projectile"
	Path              ecs.ComponentType = "path"
	PathFollow        ecs.ComponentType = "path_follow"
	Wallet            ecs.ComponentType = "wallet"
	Renderable        ecs.ComponentType = "renderable"
	ShootIntent       ecs.ComponentType = "shoot_intent"
	BuyIntent         ecs.ComponentType = "buy_intent"
	CreateTowerIntent ecs.ComponentType = "create_tower_intent"
)

type DisplayComponent struct {
	ecs.Component
	Width, Height int
}

func (c DisplayComponent) GetType() ecs.ComponentType {
	return Display
}

type CursorComponent struct {
	ecs.Component
}

func (c CursorComponent) GetType() ecs.ComponentType {
	return Cursor
}

type GameStateComponent struct {
	ecs.Component
	GameOver bool
}

func (c GameStateComponent) GetType() ecs.ComponentType {
	return GameState
}

type PlayerComponent struct {
	ecs.Component
}

func (c PlayerComponent) GetType() ecs.ComponentType {
	return Player
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

// Bounding box pairs with the Position component to define the size of an entity. Position is the center of the bounding box
type BoundingBoxComponent struct {
	ecs.Component
	Width, Height float64
}

func (c BoundingBoxComponent) GetType() ecs.ComponentType {
	return BoundingBox
}

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

type TowerComponent struct {
	ecs.Component
	Cooldown      time.Duration
	LastFired     time.Time
	Damage, Range float64
}

func (c TowerComponent) GetType() ecs.ComponentType {
	return Tower
}

type TowerType string

const (
	BasicTower  TowerType = "basic"
	MediumTower TowerType = "medium"
	HeavyTower  TowerType = "heavy"
)

type TowerTemplateComponent struct {
	ecs.Component
	Type TowerType
	Cost float64
}

func (c TowerTemplateComponent) GetType() ecs.ComponentType {
	return TowerTemplate
}

type ProjectileComponent struct {
	ecs.Component
	TargetEntity  ecs.Entity
	Damage, Speed float64
}

func (c ProjectileComponent) GetType() ecs.ComponentType {
	return Projectile
}

type PathComponent struct {
	ecs.Component
	ID        string
	Waypoints []PositionComponent
}

func (c PathComponent) GetType() ecs.ComponentType {
	return Path
}

type PathFollowComponent struct {
	ecs.Component
	PathID        string // ID of the path to follow
	WaypointIndex int    // Current waypoint
}

func (c PathFollowComponent) GetType() ecs.ComponentType {
	return PathFollow
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

type ShootIntentComponent struct {
	ecs.Component
	Shooter, Target ecs.Entity
}

func (c ShootIntentComponent) GetType() ecs.ComponentType {
	return ShootIntent
}

type BuyIntentComponent struct {
	ecs.Component
	// We can use this for upgrades later
}

func (c BuyIntentComponent) GetType() ecs.ComponentType {
	return BuyIntent
}

type CreateTowerIntentComponent struct {
	ecs.Component
	TowerType TowerType
	Position  PositionComponent
}

func (c CreateTowerIntentComponent) GetType() ecs.ComponentType {
	return CreateTowerIntent
}

var ComponentTypes = []ecs.ComponentType{
	Display,
	Cursor,
	GameState,
	Player,
	Enemy,
	BoundingBox,
	Position,
	Health,
	Velocity,
	Tower,
	TowerTemplate,
	Projectile,
	Path,
	PathFollow,
	Wallet,
	Renderable,
	ShootIntent,
	BuyIntent,
	CreateTowerIntent,
}
