package teaui

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/internal/input"
	"ecstemplate/pkg/ecs"
)

type InputManager struct {
	state                  input.InputState
	keysBuffer             []string
	minX, minY, maxX, maxY int
}

func (im *InputManager) Initialize() error {
	im.state = input.InputState{
		Actions:      make(map[input.Action]bool),
		CursorX:      10,
		CursorY:      10,
		IsPlacing:    false,
		PlacingTower: "",
	}
	im.keysBuffer = make([]string, 0)

	return nil
}

func (im *InputManager) QueueKey(key string) {
	im.keysBuffer = append(im.keysBuffer, key)
}

func (im *InputManager) Update() {
	// Reset actions
	for action := range im.state.Actions {
		im.state.Actions[action] = false
	}

	// Process queued keys
	for _, key := range im.keysBuffer {
		switch key {
		case "w", "up":
			im.state.Actions[input.ActionMoveUp] = true
			im.state.CursorY = max(im.minY, im.state.CursorY-1)
		case "s", "down":
			im.state.Actions[input.ActionMoveDown] = true
			im.state.CursorY = min(im.maxY, im.state.CursorY+1)
		case "a", "left":
			im.state.Actions[input.ActionMoveLeft] = true
			im.state.CursorX = max(im.minX, im.state.CursorX-1)
		case "d", "right":
			im.state.Actions[input.ActionMoveRight] = true
			im.state.CursorX = min(im.maxX, im.state.CursorX+1)
		case "1":
			im.state.Actions[input.ActionBuildBasic] = true
			im.state.PlacingTower = components.BasicTower
			im.state.IsPlacing = true
		case "2":
			im.state.Actions[input.ActionBuildMedium] = true
			im.state.PlacingTower = components.MediumTower
			im.state.IsPlacing = true
		case "3":
			im.state.Actions[input.ActionBuildHeavy] = true
			im.state.PlacingTower = components.HeavyTower
			im.state.IsPlacing = true
		case "enter", " ":
			im.state.Actions[input.ActionSelect] = true
		case "esc":
			im.state.Actions[input.ActionCancel] = true
			im.state.IsPlacing = false
		case "n":
			im.state.Actions[input.ActionNextWave] = true
		case "p":
			im.state.Actions[input.ActionTogglePause] = true
		case "q":
			im.state.Actions[input.ActionQuit] = true
		}
	}

	// Clear the buffer after processing
	im.keysBuffer = im.keysBuffer[:0]
}

func (im *InputManager) GetState() input.InputState {
	return im.state
}

func (im *InputManager) ProcessInputs(
	world *ecs.World,
	componentAccess *components.ComponentAccess,
) {
	if im.state.Actions[input.ActionSelect] && im.state.IsPlacing {
		// Create a tower at cursor position
		posComponent := components.PositionComponent{
			X: float64(im.state.CursorX),
			Y: float64(im.state.CursorY),
		}

		// Create intent to build tower
		createTowerIntent := &components.CreateTowerIntentComponent{
			TowerType: im.state.PlacingTower,
			Position:  posComponent,
		}

		// Find player entity to attach the intent to
		playerEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Player)
		if len(playerEnts) > 0 {
			world.ComponentManager.AddComponent(
				playerEnts[0],
				components.CreateTowerIntent,
				createTowerIntent,
			)
		}

		// Reset placement mode
		im.state.IsPlacing = false
	}
}

func (im *InputManager) SetCursorBounds(minX, minY, maxX, maxY int) {
	im.minX = minX
	im.minY = minY
	im.maxX = maxX
	im.maxY = maxY
}

func (im *InputManager) Shutdown() {
	// no-op, nothing to clean up
}
