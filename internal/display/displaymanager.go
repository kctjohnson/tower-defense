package display

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// DisplayManager is an interface that defines the methods that a display manager should implement
type DisplayManager interface {
	// Initialize sets up the display with the given dimensions
	Initialize(width, height int) error

	// Clear resets the display for the next frame
	Clear()

	// Render renders the current game state to the display
	Render(world *ecs.World, componentAccess *components.ComponentAccess)

	// RenderEntity renders a specific entity
	RenderEntity(
		entity ecs.Entity,
		position *components.PositionComponent,
		renderable *components.RenderableComponent,
	)

	// RenderUI renders UI elements like health, money, wave info
	RenderUI(gameInfo GameInfo)

	// Update refreshes the display after all elements have been rendered
	Update()

	// Shutdown cleans up resources used by the display
	Shutdown()
}

// GameInfo contains the information needed to render the UI
type GameInfo struct {
	PlayerHealth float64
	PlayerMoney  float64
	CurrentWave  int
	WaveProgress float64
	GameOver     bool
	Message      string
}
