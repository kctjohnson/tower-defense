package events

import "ecstemplate/pkg/ecs"

const (
	EnemyReachedEnd ecs.EventType = "enemy_reached_end"
	GameOver        ecs.EventType = "game_over"
)

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
