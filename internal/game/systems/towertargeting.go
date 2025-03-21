package systems

import "ecstemplate/pkg/ecs"

// TowerTargetingSystem is a system that monitors the closest enemy to each tower
// and adds a shoot intent to the tower when the wait duration has passed and a target is in range
type TowerTargetingSystem struct{}

func (s *TowerTargetingSystem) Update(world *ecs.World, deltaTime float64) {

}
