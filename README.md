# Ecosystem Simulation (Grass, Rabbits, Foxes) – Go + Ebiten + Pixel Art

This is a graphical ecosystem simulation written in Go, using the [Ebiten](https://ebiten.org/) game library. The simulation models the interaction between grass, rabbits, and foxes on a grid-based map with pixel-art-style graphics.

## Features

- **Grass** grows on tiles over time and spreads to neighboring empty spaces.
- **Rabbits** move around, eat grass, reproduce, and die of starvation if they can't eat.
- **Foxes** hunt rabbits, reproduce after eating, and also die without food.
- **Real-time graph** displaying population dynamics of all entities.
- **Tile-based pixel art** visuals (32x32 JPG/PNG sprites).

## TODO

- **Mouse or keyboard input** to spawn new rabbits and foxes on the map.
- **UI panel** to adjust simulation parameters like grass growth speed, initial populations, etc.
- **Buttons** for pausing or restarting the simulation.

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

## Running the Simulation

```bash
go run main.go
```
