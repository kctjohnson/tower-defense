package game

import (
	"log"
	"os"
	"time"

	"ecstemplate/internal/display"
	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/internal/game/systems"
	"ecstemplate/internal/game/ui/teaui"
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
	world.AddSystem(&systems.ProjectileSystem{
		ComponentAccess: componentAccess,
	})
	world.AddSystem(&systems.ProjectileCreationSystem{
		ComponentAccess: componentAccess,
	})
	world.AddSystem(&systems.TowerTargetingSystem{
		ComponentAccess: componentAccess,
	})
	world.AddSystem(&systems.CollisionSystem{
		ComponentAccess: componentAccess,
	})
	world.AddSystem(&systems.TowerFactorySystem{
		ComponentAccess: componentAccess,
	})

	return &Game{
		world:           world,
		inputManager:    &teaui.InputManager{},
		displayManager:  &teaui.DisplayManager{},
		componentAccess: componentAccess,
	}
}

func (g *Game) Initialize(width, height int) {
	// Initialize the display and input managers
	if err := g.displayManager.Initialize(width, height); err != nil {
		g.world.Logger.Fatalf("Failed to initialize display manager: %v", err)
	}

	if err := g.inputManager.Initialize(); err != nil {
		g.world.Logger.Fatalf("Failed to initialize input manager: %v", err)
	}

	// Register component types
	g.registerComponentTypes()

	// Register event handlers
	g.world.RegisterEventHandler(events.TowerCreated, g.towerCreatedEventHandler)
	g.world.RegisterEventHandler(events.EnemyKilled, g.enemyKilledEventHandler)
	g.world.RegisterEventHandler(events.ProjectileFired, g.projectileFiredEventHandler)
	g.world.RegisterEventHandler(events.EnemyReachedEnd, g.enemyReachedEndEventHandler)
	g.world.RegisterEventHandler(events.GameOver, g.gameOverEventHandler)

	// Create the player
	playerEnt := g.world.EntityManager.CreateEntity()
	g.world.ComponentManager.AddComponent(
		playerEnt,
		components.Player,
		&components.PlayerComponent{},
	)
	g.world.ComponentManager.AddComponent(
		playerEnt,
		components.Health,
		&components.HealthComponent{
			Current: 50,
			Max:     50,
		},
	)
	g.world.ComponentManager.AddComponent(
		playerEnt,
		components.Wallet,
		&components.WalletComponent{
			Money: 0,
		},
	)

	// Create the path
	pathEnt := g.world.EntityManager.CreateEntity()
	g.world.ComponentManager.AddComponent(
		pathEnt,
		components.Path,
		&components.PathComponent{
			ID: "starting-path",
			Waypoints: []components.PositionComponent{
				{X: 5, Y: 5},
				{X: 15, Y: 5},
				{X: 15, Y: 15},
			},
		},
	)

	// Create an enemy on the path
	enemyEnt := g.world.EntityManager.CreateEntity()
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.Enemy,
		&components.EnemyComponent{
			Type:   "basic",
			Speed:  1,
			Reward: 10,
		},
	)
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.BoundingBox,
		&components.BoundingBoxComponent{
			Width:  1,
			Height: 1,
		},
	)
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.Position,
		&components.PositionComponent{
			X: 5,
			Y: 5,
		},
	)
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.Health,
		&components.HealthComponent{
			Current: 10,
			Max:     10,
		},
	)
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.PathFollow,
		&components.PathFollowComponent{
			PathID:        "starting-path",
			WaypointIndex: 0,
		},
	)
	g.world.ComponentManager.AddComponent(
		enemyEnt,
		components.Renderable,
		&components.RenderableComponent{
			Symbol: "E",
		},
	)

	// Create a tower
}

func (g *Game) registerComponentTypes() {
	// Register all component types with the component manager
	for _, componentType := range components.ComponentTypes {
		g.world.ComponentManager.RegisterComponentType(componentType)
	}
}

func (g *Game) Update(deltaTime float64) {
	// Gather and process input
	g.inputManager.Update()
	g.inputManager.ProcessInputs(g.world, g.componentAccess)

	// Update the game state
	g.world.Update(deltaTime)

	// Do displaying stuff
	g.displayManager.Clear()
	g.displayManager.Render(g.world, g.componentAccess)
	g.displayManager.RenderUI(g.getGameInfo())
	g.displayManager.Update()
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

		// Update the game
		g.Update(deltaTime)

		// Sleep to maintain target frame rate
		frameTime := time.Since(frameStartTime)
		if sleepTime := targetFrameTime - frameTime; sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}

func (g *Game) GetDisplayManager() display.DisplayManager {
	return g.displayManager
}

func (g *Game) GetInputManager() input.InputManager {
	return g.inputManager
}

func (g *Game) getGameInfo() display.GameInfo {
	return display.GameInfo{
		PlayerHealth: 100,
		PlayerMoney:  100,
		CurrentWave:  1,
		WaveProgress: 0.5,
		GameOver:     false,
		Message:      "",
	}
}
