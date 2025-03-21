package events

import "ecstemplate/pkg/ecs"

const (
	TowerShot       ecs.EventType = "tower_shot"
	EnemyReachedEnd ecs.EventType = "enemy_reached_end"
	GameOver        ecs.EventType = "game_over"
)

type TowerShotEvent struct {
	Shooter ecs.Entity
	Target  ecs.Entity
}

func (e *TowerShotEvent) Type() ecs.EventType {
	return TowerShot
}

func (e *TowerShotEvent) Entity() ecs.Entity {
	return e.Shooter
}

func (e *TowerShotEvent) Data() any {
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
