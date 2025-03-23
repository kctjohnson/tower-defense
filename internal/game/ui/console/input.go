package console

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/internal/input"
	"ecstemplate/pkg/ecs"
)

// Sample Console implementation of the InputManager interface

type ConsoleInputManager struct{}

func (im *ConsoleInputManager) Initialize() error {
	return nil
}

func (im *ConsoleInputManager) Update() {

}

func (im *ConsoleInputManager) GetState() input.InputState {
	return input.InputState{}
}

func (im *ConsoleInputManager) ProcessInputs(
	world *ecs.World,
	componentAccess *components.ComponentAccess,
) {

}

func (im *ConsoleInputManager) SetCursorBounds(minX, minY, maxX, maxY int) {

}

func (im *ConsoleInputManager) Shutdown() {

}
