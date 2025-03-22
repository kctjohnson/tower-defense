package systems

import (
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// ProjectileSystem handles the movement of projectiles
type ProjectileSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *ProjectileSystem) Update(world *ecs.World, deltaTime float64) {
	// Get the screen
	displayEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Display)
	if len(displayEnts) != 1 {
		return
	}
	display, _ := s.ComponentAccess.GetDisplayComponent(displayEnts[0])

	// Get the projectiles
	projectileEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Projectile)

	// Loop through all projectiles
	for _, projectileEnt := range projectileEnts {
		projPos, _ := s.ComponentAccess.GetPositionComponent(projectileEnt)
		projVel, _ := s.ComponentAccess.GetVelocityComponent(projectileEnt)

		// Move the projectile
		projPos.X += projVel.X * deltaTime
		projPos.Y += projVel.Y * deltaTime

		// Check if the projectile is out of bounds
		if projPos.X < 0 || projPos.X > float64(display.Width) ||
			projPos.Y < 0 || projPos.Y > float64(display.Height) {
			world.ComponentManager.RemoveAllComponents(projectileEnt)
			world.EntityManager.RemoveEntity(projectileEnt)
			continue
		}
	}
}
