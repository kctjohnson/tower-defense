package ecs

import "slices"

// ComponentInterface is a marker interface for all component types
type ComponentInterface interface {
	IsComponent()
	GetType() ComponentType
}

// Component is a marker struct for all component types
type Component struct{}

func (c *Component) IsComponent() {}

type ComponentType string

// ComponentManager handles storage and retrieval of components
type ComponentManager struct {
	components map[ComponentType]map[Entity]ComponentInterface
}

func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		components: make(map[ComponentType]map[Entity]ComponentInterface),
	}
}

func (cm *ComponentManager) RegisterComponentType(componentType ComponentType) {
	if _, exists := cm.components[componentType]; !exists {
		cm.components[componentType] = make(map[Entity]ComponentInterface)
	}
}

func (cm *ComponentManager) AddComponent(
	entity Entity,
	componentType ComponentType,
	component ComponentInterface,
) {
	if _, exists := cm.components[componentType]; !exists {
		cm.RegisterComponentType(componentType)
	}
	cm.components[componentType][entity] = component
}

func (cm *ComponentManager) RemoveComponent(entity Entity, componentType ComponentType) {
	if componentMap, exists := cm.components[componentType]; exists {
		delete(componentMap, entity)
	}
}

func (cm *ComponentManager) GetComponent(
	entity Entity,
	componentType ComponentType,
) (ComponentInterface, bool) {
	if componentMap, exists := cm.components[componentType]; exists {
		component, found := componentMap[entity]
		return component, found
	}
	return nil, false
}

func (cm *ComponentManager) HasComponent(entity Entity, componentType ComponentType) bool {
	if componentMap, exists := cm.components[componentType]; exists {
		_, found := componentMap[entity]
		return found
	}
	return false
}

func (cm *ComponentManager) GetAllEntitiesWithComponent(componentType ComponentType) []Entity {
	if componentMap, exists := cm.components[componentType]; exists {
		entities := make([]Entity, 0, len(componentMap))
		for e := range componentMap {
			entities = append(entities, e)
		}
		return entities
	}
	return []Entity{}
}

// Gets entities that have all of the specified components
func (cm *ComponentManager) GetAllEntitiesWithComponents(componentTypes []ComponentType) []Entity {
	entities := make(map[Entity]map[ComponentType]bool)
	for _, componentType := range componentTypes {
		for entity := range cm.components[componentType] {
			if _, exists := entities[entity]; !exists {
				entities[entity] = make(map[ComponentType]bool)
			}
			entities[entity][componentType] = true
		}
	}

	entityList := make([]Entity, 0, len(entities))
	for entity, components := range entities {
		// Make sure the entity has the right number of components
		if len(components) == len(componentTypes) {
			// Check if the entity has all of the specified components
			valid := true
			for comp := range components {
				if !slices.Contains(componentTypes, comp) {
					valid = false
					break
				}
			}

			// Add valid entities to the list
			if valid {
				entityList = append(entityList, entity)
			}
		}
	}

	return entityList
}

func (cm *ComponentManager) RemoveAllComponents(entity Entity) {
	for _, componentMap := range cm.components {
		delete(componentMap, entity)
	}
}
