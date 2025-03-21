package systems

import (
	"log"
	"math"
	"testing"

	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

func TestEnemyMovementSystem(t *testing.T) {
	logger := log.New(log.Writer(), "TestEnemyMovementSystem: ", log.Flags())
	world := ecs.NewWorld(logger)

	// Register component types
	for _, componentType := range components.ComponentTypes {
		world.ComponentManager.RegisterComponentType(componentType)
	}

	componentAccess := components.NewComponentAccess(world)

	// Create test path
	pathEnt := world.EntityManager.CreateEntity()
	pathComponent := &components.PathComponent{
		ID: "test-path",
		Waypoints: []components.PositionComponent{
			{X: 0, Y: 0},
			{X: 10, Y: 0},
			{X: 10, Y: 10},
		},
	}
	world.ComponentManager.AddComponent(pathEnt, components.Path, pathComponent)

	// Create test enemy
	enemyEnt := world.EntityManager.CreateEntity()
	enemyComponent := &components.EnemyComponent{
		Type:   "basic",
		Speed:  1,
		Reward: 0,
	}
	positionComponent := &components.PositionComponent{
		X: 0,
		Y: 0,
	}
	healthComponent := &components.HealthComponent{
		Current: 20,
		Max:     20,
	}
	pathFollowComponent := &components.PathFollowComponent{
		PathID:        "test-path",
		WaypointIndex: 0,
	}
	renderableComponent := &components.RenderableComponent{
		Symbol: "@",
	}
	world.ComponentManager.AddComponent(enemyEnt, components.Enemy, enemyComponent)
	world.ComponentManager.AddComponent(enemyEnt, components.Position, positionComponent)
	world.ComponentManager.AddComponent(enemyEnt, components.Health, healthComponent)
	world.ComponentManager.AddComponent(enemyEnt, components.PathFollow, pathFollowComponent)
	world.ComponentManager.AddComponent(enemyEnt, components.Renderable, renderableComponent)

	// Create the system
	system := &EnemyMovementSystem{
		ComponentAccess: componentAccess,
	}

	RunSimulation(system, world, 18, 60.0)

	logger.Printf("Waypoint Index: %d", pathFollowComponent.WaypointIndex)

	// Check the enemy's position
	position, _ := componentAccess.GetPositionComponent(enemyEnt)
	expectedPosition := components.PositionComponent{X: 10, Y: 8}

	// Normalize the position to 2 decimal places
	position.X = float64(math.Round(position.X*100)) / 100
	position.Y = float64(math.Round(position.Y*100)) / 100

	if position.X != expectedPosition.X || position.Y != expectedPosition.Y {
		t.Errorf("Expected enemy position to be %v, got %v", expectedPosition, position)
	}
}

// Run simulation for specified number of seconds at 60 FPS
func RunSimulation(system ecs.System, world *ecs.World, seconds float64, fps float64) {
	totalFrames := int(seconds * fps)
	deltaTime := 1.0 / fps

	for range totalFrames {
		system.Update(world, deltaTime)
	}
}
