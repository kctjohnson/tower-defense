# Tower Defense Terminal Game Development Plan

This was generated using Claude as a loose guide. I've already changed a few things foundationally, so this will likely
fall out of date quickly, as I'm not going to be following it to the letter.

## Phase 1: Core Structure Setup

1. Define your component types:

   - GameState (gameover bool)
   - Player (just a marker)
   - Enemy (type, speed, reward float64)
   - Position (x, y float64)
   - Health (current, max float64)
   - Velocity (x, y float64)
   - Tower (damage, range, cooldown, lastFired float64)
   - Projectile (damage, speed, targetID Entity)
   - Path (id string, waypoints []Position)
   - PathFollow (pathID string, currentWaypoint int)
   - Wallet (money int)
   - Renderable (symbol string, color/style)
   - ShootIntent (shooter, target Entity)

1. Create game-specific systems:

   - EnemyMovementSystem (moves enemies along paths)
   - TowerTargetingSystem (towers select enemies in range and fire)
   - ProjectileCreationSystem (handles projectile creation from tower intents)
   - ProjectileSystem (creates and moves projectiles)
   - CollisionSystem (handles projectile hits)
   - EconomySystem (handles money from kills, tower purchases)
   - WaveSystem (spawns enemy waves)
   - RenderSystem (displays the game state in terminal)

1. Define events:

   - EnemySpawned
   - EnemyReachedEnd
   - EnemyKilled
   - TowerBuilt
   - ProjectileFired
   - WaveStarted
   - WaveCompleted
   - GameOver

## Phase 2: Game Setup and Map

1. Create a map representation with:

   - Path for enemies
   - Valid tower placement spots
   - Start and end positions

1. Implement path finding for enemies:

   - Define waypoints along the path
   - Add logic for enemies to follow waypoints

1. Implement the game initialization:

   - Create initial entities (path, wallet)
   - Setup the wave system
   - Configure game parameters (starting money, enemy stats, tower stats)

## Phase 3: Enemy System

1. Implement the enemy spawning logic:

   - Create enemy factory function
   - Set up wave patterns (number, type, timing)

1. Implement enemy movement system:

   - Move enemies along path waypoints
   - Handle different enemy speeds
   - Detect when enemies reach the end

1. Add enemy health and damage handling:

   - Implement health reduction on projectile hits
   - Handle enemy death and reward

## Phase 4: Tower System

1. Implement tower placement:

   - Validate placement locations
   - Handle tower purchase and placement

1. Implement tower targeting:

   - Search for enemies in range
   - Target selection strategy (first, strongest, weakest, closest)

1. Implement tower firing:

   - Handle cooldown timers
   - Spawn projectiles
   - Different tower types (if desired)

## Phase 5: Projectile and Combat System

1. Implement projectile movement:

   - Update positions based on velocity
   - Handle out-of-bounds cleanup

1. Implement collision detection:

   - Check projectile-enemy overlaps
   - Apply damage when collisions occur
   - Remove projectiles after hits

1. Implement enemy destruction:

   - Update wallet on enemy kill
   - Clean up enemy entities
   - Award player

## Phase 6: Game Flow Systems

1. Implement wave management:

   - Control timing between waves
   - Increase difficulty over time
   - Track wave progress

1. Implement economy system:

   - Track player money
   - Handle tower purchasing
   - Potential upgrades

1. Add win/lose conditions:

   - Detect when player has lost (too many enemies reached end)
   - Potential win condition (survive all waves)

## Phase 7: Terminal Rendering

1. Create a terminal display system:

   - Convert floating-point positions to terminal grid
   - Choose symbols for different entity types
   - Use colors/styles if your terminal supports them

1. Add UI elements:

   - Status display (money, health, wave)
   - Tower selection interface
   - Messages for events

1. Implement input handling:

   - Tower placement
   - Tower selection
   - Game control (pause, restart)

## Phase 8: Polishing

1. Add game balance:

   - Tune enemy health, speed, rewards
   - Adjust tower costs, damage, range
   - Create interesting wave patterns

1. Add feedback effects:

   - Visual indicators for hits
   - Messages for important events
   - Clear indication of game state

1. Performance optimization:

   - Efficient entity queries
   - Collision optimizations
   - Rendering optimizations

1. Add save/load capability if desired
