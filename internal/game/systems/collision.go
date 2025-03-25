package systems

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/pkg/ecs"
)

// CollisionSystem checks for collisions between projectiles and enemies, and handles them accordingly
type CollisionSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *CollisionSystem) Update(world *ecs.World, deltaTime float64) {
	// Get the projectiles
	projectileEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Projectile)

	// Get the enemies
	enemyEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Enemy)

	// Loop through all projectiles
	for _, projectileEnt := range projectileEnts {
		projPos, h1 := s.ComponentAccess.GetPositionComponent(projectileEnt)
		projBoundingBox, h2 := s.ComponentAccess.GetBoundingBoxComponent(projectileEnt)

		if h1 == false || h2 == false {
			continue
		}

		// Loop through all enemies
		for _, enemyEnt := range enemyEnts {
			enemyPos, h3 := s.ComponentAccess.GetPositionComponent(enemyEnt)
			enemyBoundingBox, h4 := s.ComponentAccess.GetBoundingBoxComponent(enemyEnt)

			if h3 == false || h4 == false {
				continue
			}

			// Check if the projectile is colliding with the enemy
			if isColliding(*projPos, *enemyPos, *projBoundingBox, *enemyBoundingBox) {
				proj, _ := s.ComponentAccess.GetProjectileComponent(projectileEnt)
				enemyHealth, _ := s.ComponentAccess.GetHealthComponent(enemyEnt)

				// Remove the projectile
				world.ComponentManager.RemoveAllComponents(projectileEnt)
				world.EntityManager.RemoveEntity(projectileEnt)

				// Decrease the enemy health
				enemyHealth.Current -= proj.Damage

				// Check if the enemy is dead
				if enemyHealth.Current <= 0 {
					// Get the player and wallet components
					playerEnts := world.ComponentManager.GetAllEntitiesWithComponent(
						components.Player,
					)
					if len(playerEnts) != 1 {
						continue
					}
					playerEnt := playerEnts[0]
					wallet, _ := s.ComponentAccess.GetWalletComponent(playerEnt)

					enemy, _ := s.ComponentAccess.GetEnemyComponent(enemyEnt)
					wallet.Money += enemy.Reward

					// Queue enemy killed event
					world.QueueEvent(&events.EnemyKilledEvent{
						EnemyType: "enemy",
						Reward:    enemy.Reward,
					})

					// Remove the enemy
					world.ComponentManager.RemoveAllComponents(enemyEnt)
					world.EntityManager.RemoveEntity(enemyEnt)
				}
				break
			}
		}
	}
}

func isColliding(
	aPos, bPos components.PositionComponent,
	aBox, bBox components.BoundingBoxComponent,
) bool {
	// Position is in the center of the bounding box
	aLeft := aPos.X - float64(aBox.Width)/2
	aRight := aPos.X + float64(aBox.Width)/2
	aTop := aPos.Y - float64(aBox.Height)/2
	aBottom := aPos.Y + float64(aBox.Height)/2

	bLeft := bPos.X - float64(bBox.Width)/2
	bRight := bPos.X + float64(bBox.Width)/2
	bTop := bPos.Y - float64(bBox.Height)/2
	bBottom := bPos.Y + float64(bBox.Height)/2

	return aRight >= bLeft && aLeft <= bRight && aBottom >= bTop && aTop <= bBottom
}
