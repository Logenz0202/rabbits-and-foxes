package main

import (
	"math/rand"
	"time"

	"rabbits-and-foxes/internal/graphics"
	"rabbits-and-foxes/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *world.World
	tick  int
}

func (g *Game) Update() error {
	if g.tick%5 == 0 {
		g.world.Update()
	}
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	graphics.RenderWorld(screen, g.world)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return world.MapWidth * graphics.TileSize, world.MapHeight * graphics.TileSize
}

func main() {
	rand.Seed(time.Now().UnixNano())

	graphics.LoadAssets()
	w := world.NewWorld()

	for y := 0; y < world.MapHeight; y++ {
		for x := 0; x < world.MapWidth; x++ {
			if rand.Float64() < 0.05 {
				w.Tiles[y][x].Grass = world.NewGrass()
			}
		}
	}

	for i := 0; i < 20; i++ {
		x := rand.Intn(world.MapWidth)
		y := rand.Intn(world.MapHeight)
		w.Rabbits = append(w.Rabbits, world.NewRabbit(x, y))
	}

	ebiten.SetWindowTitle("Rabbits and Foxes – Ecosystem")
	ebiten.SetWindowSize(800, 800)

	if err := ebiten.RunGame(&Game{world: w}); err != nil {
		panic(err)
	}
}
