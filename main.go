package main

import (
	"math/rand"
	"rabbits-and-foxes/internal/game"
	"rabbits-and-foxes/internal/graphics"
	"rabbits-and-foxes/internal/world"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	graphics.LoadAssets()
	g := game.NewGame()

	for y := 0; y < world.MapHeight; y++ {
		for x := 0; x < world.MapWidth; x++ {
			if rand.Float64() < world.InitialGrassDensity {
				g.World.Tiles[y][x].Grass = world.NewGrass()
			}
		}
	}

	for i := 0; i < world.InitialRabbitCount; i++ {
		x := rand.Intn(world.MapWidth)
		y := rand.Intn(world.MapHeight)
		g.World.Rabbits = append(g.World.Rabbits, world.NewRabbit(x, y))
	}

	for i := 0; i < world.InitialFoxCount; i++ {
		x := rand.Intn(world.MapWidth)
		y := rand.Intn(world.MapHeight)
		g.World.Foxes = append(g.World.Foxes, world.NewFox(x, y))
	}

	ebiten.SetWindowTitle("Rabbits and Foxes – Ecosystem")
	ebiten.SetWindowSize(956, 800)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
