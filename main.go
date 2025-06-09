package main

import (
	"math/rand/v2"
	"rabbits-and-foxes/internal/game"
	"rabbits-and-foxes/internal/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *game.World
}

func (g *Game) Update() error {
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	graphics.RenderWorld(screen, g.world)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return game.MapWidth * graphics.TileSize, game.MapHeight * graphics.TileSize
}

func main() {
	graphics.LoadAssets()
	world := game.NewWorld()

	for y := 0; y < game.MapHeight; y++ {
		for x := 0; x < game.MapWidth; x++ {
			if rand.Float64() < 0.05 {
				world.Tiles[y][x].Grass = game.NewGrass()
			}
		}
	}

	ebiten.SetWindowTitle("Symulacja – trawa i mapa")
	ebiten.SetWindowSize(800, 800)

	if err := ebiten.RunGame(&Game{world: world}); err != nil {
		panic(err)
	}
}
