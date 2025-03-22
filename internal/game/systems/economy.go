package systems

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// EconomySystem handles buy intents from player for upgrades
type EconomySystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *EconomySystem) Update(world *ecs.World, deltaTime float64) {
	// Get all entities with BuyIntent component
}
