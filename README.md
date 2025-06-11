# Ecosystem Simulation (Grass, Rabbits, Foxes) – Go + Ebiten + Pixel Art

This is a graphical ecosystem simulation written in Go, using the [Ebiten](https://ebiten.org/) game library. The simulation models the interaction between grass, rabbits, and foxes on a grid-based map with pixel-art-style graphics.

## Documentation ([PL](dokumentacja.md)) / ([EN](documentation.md))

## Features

- **Grass** grows on tiles over time and spreads to neighboring empty spaces.
- **Rabbits** move around, eat grass, reproduce, and die of starvation if they can't eat.
- **Foxes** hunt rabbits, reproduce after eating, and also die without food.
- **Real-time graph** displaying population dynamics of all entities.
- **Tile-based pixel art** visuals (32x32 JPG/PNG sprites).
- **Keyboard input** for pausing and restarting the simulation.

## TODO

- **Mouse or keyboard input** to spawn new rabbits and foxes on the map.
- **UI panel** to adjust simulation parameters like grass growth speed, initial populations, etc.

## Project Structure

```
project_root/
├── internal/
│   │
│   ├── graphics/           # Asset loading and rendering code
│   │   ├── assets.go
│   │   └── assets/
│   │       ├── animals/
│   │       └── tiles/
│   │
│   ├── world/              # Map, grass logic, entities
│   └── game/               # Game loop, update logic
│
├── main.go
├── go.mod
└── README.md
```

## Requirements

- Go 1.18+
- Ebiten (via `go get github.com/hajimehoshi/ebiten/v2`)

## Changing simulation settings

All changes are currently made via `world/world.go`.

```go
const (
	MapWidth            = 32   // number of vertical tiles
	MapHeight           = 32   // number of horizontal tiles
	InitialGrassDensity = 0.15 // % of tiles with pre-planted grass
	InitialRabbitCount  = 100  // initial number of rabbits
	InitialFoxCount     = 20   // initial number of foxes
	SimulationSpeed     = 7    // the higher the number, the slower the simulation
)
```

## Running the Simulation

```bash
go run main.go
```
