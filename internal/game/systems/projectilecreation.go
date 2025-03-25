package systems

import (
	"math"

	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

// ProjectileCreationSystem handles the shootintents and creates the projectiles accordingly
type ProjectileCreationSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *ProjectileCreationSystem) Update(world *ecs.World, deltaTime float64) {
	// Get all entities with a shoot intent (should only be towers)
	shootIntentEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.ShootIntent)

	// Loop through all shoot intents
	for _, shootIntentEnt := range shootIntentEnts {
		shootIntent, _ := s.ComponentAccess.GetShootIntentComponent(shootIntentEnt)

		// Get the shooter's position
		shooterPos, _ := s.ComponentAccess.GetPositionComponent(shootIntent.Shooter)

		// Get the target's position
		targetPos, _ := s.ComponentAccess.GetPositionComponent(shootIntent.Target)

		// Get the angle between the two
		angle := calcAngleBetweenPoints(*shooterPos, *targetPos)

		// Create the projectile entity with the velocity vector
		baseProjectileSpeed := 2.0 // This should be a constant
		projectileEnt := world.EntityManager.CreateEntity()
		world.ComponentManager.AddComponent(
			projectileEnt,
			components.Projectile,
			&components.ProjectileComponent{
				Damage:       1,
				Speed:        baseProjectileSpeed,
				TargetEntity: shootIntent.Target,
			},
		)
		world.ComponentManager.AddComponent(
			projectileEnt,
			components.Position,
			&components.PositionComponent{
				X: shooterPos.X,
				Y: shooterPos.Y,
			},
		)
		world.ComponentManager.AddComponent(
			projectileEnt,
			components.BoundingBox,
			&components.BoundingBoxComponent{
				Width:  1,
				Height: 1,
			},
		)
		world.ComponentManager.AddComponent(
			projectileEnt,
			components.Velocity,
			&components.VelocityComponent{
				X: math.Cos(angle) * baseProjectileSpeed,
				Y: math.Sin(angle) * baseProjectileSpeed,
			},
		)
		world.ComponentManager.AddComponent(
			projectileEnt,
			components.Renderable,
			&components.RenderableComponent{
				Symbol: "o",
			},
		)

		// Remove the shoot intent component from the tower
		world.ComponentManager.RemoveComponent(shootIntentEnt, components.ShootIntent)
	}
}
