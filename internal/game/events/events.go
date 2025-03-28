package events

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

const (
	TowerCreated    ecs.EventType = "tower_created"
	EnemyKilled     ecs.EventType = "enemy_killed"
	ProjectileFired ecs.EventType = "projectile_fired"
	EnemyReachedEnd ecs.EventType = "enemy_reached_end"
	GameOver        ecs.EventType = "game_over"
)

type TowerCreatedEvent struct {
	TowerType   components.TowerType
	TowerEntity ecs.Entity
}

func (e *TowerCreatedEvent) Type() ecs.EventType {
	return TowerCreated
}

func (e *TowerCreatedEvent) Entity() ecs.Entity {
	return e.TowerEntity
}

func (e *TowerCreatedEvent) Data() any {
	return map[string]any{
		"towerType":   e.TowerType,
		"towerEntity": e.TowerEntity,
	}
}

type EnemyKilledEvent struct {
	EnemyType string
	Reward    float64
}

func (e *EnemyKilledEvent) Type() ecs.EventType {
	return EnemyKilled
}

func (e *EnemyKilledEvent) Entity() ecs.Entity {
	return -1
}

func (e *EnemyKilledEvent) Data() any {
	return map[string]any{
		"enemyType": e.EnemyType,
		"reward":    e.Reward,
	}
}

type ProjectileFiredEvent struct {
	Shooter ecs.Entity
	Target  ecs.Entity
}

func (e *ProjectileFiredEvent) Type() ecs.EventType {
	return ProjectileFired
}

func (e *ProjectileFiredEvent) Entity() ecs.Entity {
	return e.Shooter
}

func (e *ProjectileFiredEvent) Data() any {
	return map[string]ecs.Entity{
		"shooter": e.Shooter,
		"target":  e.Target,
	}
}

type EnemyReachedEndEvent struct {
	Ent ecs.Entity
}

func (e *EnemyReachedEndEvent) Type() ecs.EventType {
	return EnemyReachedEnd
}

func (e *EnemyReachedEndEvent) Entity() ecs.Entity {
	return e.Ent
}

func (e *EnemyReachedEndEvent) Data() any {
	return nil
}

type GameOverEvent struct{}

func (e *GameOverEvent) Type() ecs.EventType {
	return GameOver
}

func (e *GameOverEvent) Entity() ecs.Entity {
	return -1
}

func (e *GameOverEvent) Data() any {
	return nil
}
