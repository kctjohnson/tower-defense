package systems

import (
	"time"

	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/pkg/ecs"
)

// TowerTargetingSystem is a system that monitors the closest enemy to each tower
// and adds a shoot intent to the tower when the wait duration has passed and a target is in range
type TowerTargetingSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *TowerTargetingSystem) Update(world *ecs.World, deltaTime float64) {
	// Get all towerEnts active in the world
	towerEnts := world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{components.Tower, components.Position, components.Renderable},
	)

	// Get all enemyEnts active in the world
	enemyEnts := world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{
			components.Enemy,
			components.Position,
			components.Health,
			components.PathFollow,
			components.Renderable,
		},
	)

	// Loop through all towers
	for _, towerEnt := range towerEnts {
		// Get the tower
		tower, _ := s.ComponentAccess.GetTowerComponent(towerEnt)

		// Get tower position
		towerPos, _ := s.ComponentAccess.GetPositionComponent(towerEnt)

		// Get the closest enemy to the tower
		closestEnemyEnt, found := s.getClosestEnemy(towerPos, tower.Range, enemyEnts)
		if !found {
			// No enemies in range
			continue
		}

		// If the tower has no cooldown, add a shoot intent
		if time.Since(tower.LastFired) >= tower.Cooldown {
			// Add a shoot intent to the tower
			world.ComponentManager.AddComponent(
				towerEnt,
				components.ShootIntent,
				&components.ShootIntentComponent{
					Shooter: towerEnt,
					Target:  closestEnemyEnt,
				},
			)

			// Queue the tower shot event
			world.QueueEvent(&events.TowerShotEvent{
				Shooter: towerEnt,
				Target:  closestEnemyEnt,
			})

			// Reset the last fired time
			tower.LastFired = time.Now()
		}
	}
}

func (s *TowerTargetingSystem) getClosestEnemy(
	towerPos *components.PositionComponent,
	towerRange float64,
	enemyEnts []ecs.Entity,
) (closest ecs.Entity, found bool) {
	closest = -1
	closestDist := 999999.0

	if len(enemyEnts) == 0 {
		return -1, false
	}

	for _, enemyEnt := range enemyEnts {
		enemyPos, _ := s.ComponentAccess.GetPositionComponent(enemyEnt)
		dist := distance(*towerPos, *enemyPos)
		if dist < closestDist && dist <= towerRange {
			closestDist = dist
			closest = enemyEnt
		}
	}

	if closest == -1 {
		return -1, false
	}

	return closest, true
}
