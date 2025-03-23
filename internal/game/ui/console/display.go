package console

import (
	"ecstemplate/internal/display"
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// Sample Console implementation of the DisplayManager interface

type ConsoleDisplayManager struct {
	width, height int
	buffer        [][]rune
}

func (dm *ConsoleDisplayManager) Initialize(width, height int) error {
	return nil
}

func (dm *ConsoleDisplayManager) Clear() {

}

func (dm *ConsoleDisplayManager) Render(
	world *ecs.World,
	componentAccess *components.ComponentAccess,
) {

}

func (dm *ConsoleDisplayManager) RenderEntity(
	entity ecs.Entity,
	position *components.PositionComponent,
	renderable *components.RenderableComponent,
) {

}

func (dm *ConsoleDisplayManager) RenderUI(gameInfo display.GameInfo) {

}

func (dm *ConsoleDisplayManager) Update() {

}

func (dm *ConsoleDisplayManager) Shutdown() {

}
