package game

import (
	"log"
	"os"
	"time"

	"ecstemplate/internal/display"
	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/internal/game/systems"
	"ecstemplate/internal/game/ui/console"
	"ecstemplate/internal/input"
	"ecstemplate/pkg/ecs"
)

type Game struct {
	world           *ecs.World
	inputManager    input.InputManager
	displayManager  display.DisplayManager
	componentAccess *components.ComponentAccess
}

func NewGame() *Game {
	logger := log.New(os.Stdout, "Game: ", log.LstdFlags)

	world := ecs.NewWorld(logger)

	// Create the component access manager
	componentAccess := components.NewComponentAccess(world)

	// Register core ECS systems
	world.AddSystem(&systems.EnemyMovementSystem{
		ComponentAccess: componentAccess,
	})

	return &Game{
		world:           world,
		inputManager:    &console.ConsoleInputManager{},
		displayManager:  &console.ConsoleDisplayManager{},
		componentAccess: componentAccess,
	}
}

func (g *Game) Initialize() {
	// Register component types
	g.registerComponentTypes()

	// Register event handlers
	g.world.RegisterEventHandler(events.Sample, g.sampleEventHandler)

	// Create the entities
}

func (g *Game) registerComponentTypes() {
	// Register all component types with the component manager
	for _, componentType := range components.ComponentTypes {
		g.world.ComponentManager.RegisterComponentType(componentType)
	}
}

func (g *Game) Run() {
	g.world.Logger.Println("Starting game...")

	targetFrameTime := time.Second / 60
	lastUpdatedTime := time.Now()

	// Main game loop
	for {
		// Start timing this fram
		frameStartTime := time.Now()

		// Calculate delta time
		deltaTime := frameStartTime.Sub(lastUpdatedTime).Seconds()
		lastUpdatedTime = frameStartTime

		// Do displaying stuff

		// Gather input

		// Update the game state
		g.world.Update(deltaTime)

		// Sleep to maintain target frame rate
		frameTime := time.Since(frameStartTime)
		if sleepTime := targetFrameTime - frameTime; sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
