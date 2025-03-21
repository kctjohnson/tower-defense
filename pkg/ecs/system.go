package ecs

// System processes entities with specific components
type System interface {
	Update(world *World, deltaTime float64)
}
