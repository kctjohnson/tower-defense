package systems

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// SampleSystem doesn't really do anything, but it could
type SampleSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *SampleSystem) Update(world *ecs.World) {
	// Manage components, states, intents, etc.
	// Queue events for the event handlers
}
