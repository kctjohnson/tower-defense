package systems

import (
	"time"

	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// WaveSystem handles the spawning of enemies when the wave timer has passed
type WaveSystem struct {
	ComponentAccess *components.ComponentAccess
	lastSpawnTime   time.Time
	cooldown        time.Duration
}

func NewWaveSystem(
	componentAccess *components.ComponentAccess,
	cooldown time.Duration,
) *WaveSystem {
	return &WaveSystem{
		ComponentAccess: componentAccess,
		cooldown:        cooldown,
		lastSpawnTime:   time.Now(),
	}
}

func (s *WaveSystem) Update(world *ecs.World, deltaTime float64) {
	if time.Since(s.lastSpawnTime) >= s.cooldown {
		// Spawn a new enemy at the start of the first path
		pathEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Path)
		if len(pathEnts) == 0 {
			return
		}
		path, _ := s.ComponentAccess.GetPathComponent(pathEnts[0])

		// Create the enemy entity
		enemyEnt := world.EntityManager.CreateEntity()
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.Enemy,
			&components.EnemyComponent{
				Type:   "basic",
				Speed:  1,
				Reward: 10,
			},
		)
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.BoundingBox,
			&components.BoundingBoxComponent{
				Width:  1,
				Height: 1,
			},
		)
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.Position,
			&components.PositionComponent{
				X: path.Waypoints[0].X,
				Y: path.Waypoints[0].Y,
			},
		)
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.Health,
			&components.HealthComponent{
				Current: 10,
				Max:     10,
			},
		)
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.PathFollow,
			&components.PathFollowComponent{
				PathID:        "starting-path",
				WaypointIndex: 0,
			},
		)
		world.ComponentManager.AddComponent(
			enemyEnt,
			components.Renderable,
			&components.RenderableComponent{
				Symbol: "E",
			},
		)

		// Reset the last spawn time
		s.lastSpawnTime = time.Now()
		s.cooldown -= 50 * time.Millisecond
	}
}
