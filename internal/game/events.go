package game

import (
	"fmt"

	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/pkg/ecs"
)

func (g *Game) towerShotEventHandler(event ecs.EventInterface) {

}

func (g *Game) enemyReachedEndEventHandler(event ecs.EventInterface) {
	// Determine the enemy damage
	enemy, _ := g.componentAccess.GetEnemyComponent(event.Entity())
	var enemyDamage float64
	switch enemy.Type {
	case "basic":
		enemyDamage = 1
	case "fast":
		enemyDamage = 2
	case "tank":
		enemyDamage = 3
	default:
		panic("Unknown enemy type")
	}

	// Get the player entity
	playerEnt := g.world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{
			components.Player,
			components.Health,
			components.Wallet,
		},
	)
	if len(playerEnt) != 1 {
		panic("Expected exactly one player entity")
	}

	health, _ := g.componentAccess.GetHealthComponent(playerEnt[0])
	health.Current -= enemyDamage
	if health.Current <= 0 {
		gameStateEnts := g.world.ComponentManager.GetAllEntitiesWithComponent(components.GameState)
		if len(gameStateEnts) != 1 {
			panic("Expected exactly one game state entity")
		}

		gameState, _ := g.componentAccess.GetGameStateComponent(gameStateEnts[0])
		gameState.GameOver = true

		g.world.QueueEvent(&events.GameOverEvent{})
	}

	// Remove the enemy entity, and all of its components
	g.world.ComponentManager.RemoveAllComponents(event.Entity())
}

func (g *Game) gameOverEventHandler(event ecs.EventInterface) {
	fmt.Println("Game Over")
	// The player has lost
	// End the game
}
