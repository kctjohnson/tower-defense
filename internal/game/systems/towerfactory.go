package systems

import (
	"errors"
	"time"

	"ecstemplate/internal/game/components"
	"ecstemplate/internal/game/events"
	"ecstemplate/pkg/ecs"
)

type TowerFactorySystem struct {
	ComponentAccess *components.ComponentAccess
	Templates       map[components.TowerType]ecs.Entity
}

func (s *TowerFactorySystem) Initialize(world *ecs.World) {
	s.Templates = make(map[components.TowerType]ecs.Entity)

	s.Templates[components.BasicTower] = world.EntityManager.CreateEntity()
	world.ComponentManager.AddComponent(
		s.Templates[components.BasicTower],
		components.TowerTemplate,
		&components.TowerTemplateComponent{
			Type: components.BasicTower,
			Cost: 5,
		},
	)
	world.ComponentManager.AddComponent(
		s.Templates[components.BasicTower],
		components.Tower,
		&components.TowerComponent{
			Cooldown:  3,
			LastFired: time.Now(),
			Damage:    1,
			Range:     5,
		},
	)

	s.Templates[components.MediumTower] = world.EntityManager.CreateEntity()
	world.ComponentManager.AddComponent(
		s.Templates[components.MediumTower],
		components.TowerTemplate,
		&components.TowerTemplateComponent{
			Type: components.MediumTower,
			Cost: 10,
		},
	)
	world.ComponentManager.AddComponent(
		s.Templates[components.MediumTower],
		components.Tower,
		&components.TowerComponent{
			Cooldown:  2,
			LastFired: time.Now(),
			Damage:    2,
			Range:     7,
		},
	)

	s.Templates[components.HeavyTower] = world.EntityManager.CreateEntity()
	world.ComponentManager.AddComponent(
		s.Templates[components.HeavyTower],
		components.TowerTemplate,
		&components.TowerTemplateComponent{
			Type: components.HeavyTower,
			Cost: 15,
		},
	)
	world.ComponentManager.AddComponent(
		s.Templates[components.HeavyTower],
		components.Tower,
		&components.TowerComponent{
			Cooldown:  1,
			LastFired: time.Now(),
			Damage:    3,
			Range:     10,
		},
	)
}

func (s *TowerFactorySystem) Update(world *ecs.World, deltaTime float64) {
	// Get the player entity
	playerEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Player)
	if len(playerEnts) != 1 {
		return
	}
	playerEnt := playerEnts[0]
	wallet, _ := s.ComponentAccess.GetWalletComponent(playerEnt)

	// Get all CreateTowerIntent components
	createTowerIntentEnts := world.ComponentManager.GetAllEntitiesWithComponent(
		components.CreateTowerIntent,
	)
	for _, createTowerIntentEnt := range createTowerIntentEnts {
		// Get the CreateTowerIntent component
		createTowerIntent, _ := s.ComponentAccess.GetCreateTowerIntentComponent(
			createTowerIntentEnt,
		)

		// Check if the player has enough money
		towerTemplate, _ := s.ComponentAccess.GetTowerTemplateComponent(
			s.Templates[createTowerIntent.TowerType],
		)

		if wallet.Money < towerTemplate.Cost {
			continue
		}

		// Create the tower at the specified position
		newTowerEnt, err := s.createTower(
			world,
			createTowerIntent.TowerType,
			createTowerIntent.Position,
		)
		if err != nil {
			return
		}

		// Deduct the cost of the tower from the player's wallet
		wallet.Money -= towerTemplate.Cost

		// Queue the new tower created event
		world.QueueEvent(&events.TowerCreatedEvent{
			TowerType:   createTowerIntent.TowerType,
			TowerEntity: newTowerEnt,
		})

		// Remove the CreateTowerIntent component
		world.ComponentManager.RemoveComponent(
			createTowerIntentEnt,
			components.CreateTowerIntent,
		)
	}
}

func (s *TowerFactorySystem) createTower(
	world *ecs.World,
	towerType components.TowerType,
	position components.PositionComponent,
) (ecs.Entity, error) {
	towerTemplateEnt, found := s.Templates[towerType]
	if !found {
		return -1, errors.New("tower type not found")
	}
	towerComp, _ := s.ComponentAccess.GetTowerComponent(towerTemplateEnt)

	tower := world.EntityManager.CreateEntity()
	world.ComponentManager.AddComponent(
		tower,
		components.Tower,
		&components.TowerComponent{
			Cooldown:  towerComp.Cooldown,
			LastFired: time.Now(),
			Damage:    towerComp.Damage,
			Range:     towerComp.Range,
		},
	)
	world.ComponentManager.AddComponent(
		tower,
		components.Position,
		&position,
	)

	return tower, nil
}
