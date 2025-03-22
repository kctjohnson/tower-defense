package input

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// Action represents player input actions
type Action string

const (
	ActionNone        Action = "none"
	ActionMoveUp      Action = "move_up"
	ActionMoveDown    Action = "move_down"
	ActionMoveLeft    Action = "move_left"
	ActionMoveRight   Action = "move_right"
	ActionSelect      Action = "select"
	ActionCancel      Action = "cancel"
	ActionBuildBasic  Action = "build_basic"
	ActionBuildMedium Action = "build_medium"
	ActionBuildHeavy  Action = "build_heavy"
	ActionNextWave    Action = "next_wave"
	ActionTogglePause Action = "toggle_pause"
	ActionQuit        Action = "quit"
)

// InputState represents the current state of all inputs
type InputState struct {
	Actions          map[Action]bool // True if action is active
	CursorX, CursorY int             // Position of cursor for placement
	PlacingTower     components.TowerType
	IsPlacing        bool
}

// InputManager is an interface that defines the methods that an input manager should implement
type InputManager interface {
	// Initialize sets up the input system
	Initialize() error

	// Update processes pending input events and updates the input state
	Update()

	// GetState returns the current input state
	GetState() InputState

	// ProcessInputs applies the current input state to the game world
	ProcessInputs(world *ecs.World, componentAccess *components.ComponentAccess)

	// SetCursorBounds sets boundaries for cursor movement
	SetCursorBounds(minX, minY, maxX, maxY int)

	// Shutdown cleans up resources
	Shutdown()
}
