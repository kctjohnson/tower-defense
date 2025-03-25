package systems

import (
	"math"

	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/pkg/ecs"
)

// EnemyMovementSystem is a system that moves enemies along a path.
type EnemyMovementSystem struct {
	ComponentAccess *components.ComponentAccess
}

func (s *EnemyMovementSystem) Update(world *ecs.World, deltaTime float64) {
	// Get all entities with an enemy, position, and path component.
	runnerEnts := world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{
			components.Position,
			components.Health,
			components.Enemy,
			components.PathFollow,
			components.Renderable,
		},
	)

	pathEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Path)
	if len(pathEnts) == 0 {
		// No paths exist in the world
		return
	}

	for _, runnerEnt := range runnerEnts {
		// Get the enemy, position, and path components for the entity.
		enemy, _ := s.ComponentAccess.GetEnemyComponent(runnerEnt)
		position, _ := s.ComponentAccess.GetPositionComponent(runnerEnt)
		pathFollow, _ := s.ComponentAccess.GetPathFollowComponent(runnerEnt)

		// Get the path for the enemy
		var path *components.PathComponent
		for _, ent := range pathEnts {
			p, _ := s.ComponentAccess.GetPathComponent(ent)
			if p.ID == pathFollow.PathID {
				path, _ = s.ComponentAccess.GetPathComponent(ent)
				break
			}
		}

		if path == nil {
			// The path for the enemy does not exist
			return
		}

		if pathFollow.WaypointIndex >= len(path.Waypoints)-1 {
			// The enemy has reached the end of the path
			world.QueueEvent(&events.EnemyReachedEndEvent{
				Ent: runnerEnt,
			})
			continue
		}

		// Get the length of the current path
		startPoint := path.Waypoints[pathFollow.WaypointIndex]
		endPoint := path.Waypoints[pathFollow.WaypointIndex+1]

		pathAngle := calcAngleBetweenPoints(startPoint, endPoint)
		distanceToMove := enemy.Speed * deltaTime
		newPosition := movePointDistance(*position, distanceToMove, pathAngle)

		distanceTraveled := distance(*position, newPosition)
		distanceToWaypoint := distance(*position, endPoint)

		// Check if the enemy has reached the end of the path
		if distanceTraveled >= distanceToWaypoint {
			pathFollow.WaypointIndex++
			if pathFollow.WaypointIndex >= len(path.Waypoints) {
				// The enemy has reached the end of the path
				world.QueueEvent(&events.EnemyReachedEndEvent{
					Ent: runnerEnt,
				})
				continue
			}

			// Position the enemy at the next waypoint
			position.X = endPoint.X
			position.Y = endPoint.Y
		} else {
			// Move the enemy along the path
			position.X = newPosition.X
			position.Y = newPosition.Y
		}
	}
}

func distance(start, end components.PositionComponent) float64 {
	// Calculate the hypotenuse
	xDiff := end.X - start.X
	yDiff := end.Y - start.Y
	return math.Sqrt((xDiff * xDiff) + (yDiff * yDiff))
}

func movePointDistance(
	point components.PositionComponent,
	distance float64,
	angle float64,
) components.PositionComponent {
	// Move the point by the given distance
	x := point.X + distance*math.Cos(angle)
	y := point.Y + distance*math.Sin(angle)
	return components.PositionComponent{X: x, Y: y}
}

func calcAngleBetweenPoints(
	start, end components.PositionComponent,
) float64 {
	// Calculate the angle between two points
	xDiff := end.X - start.X
	yDiff := end.Y - start.Y
	return math.Atan2(yDiff, xDiff)
}
