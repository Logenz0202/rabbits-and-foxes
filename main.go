package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"rabbits-and-foxes/internal/graphics"
	"rabbits-and-foxes/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	world             *world.World
	tick              int
	populationHistory struct {
		rabbits []int
		foxes   []int
		maxLen  int
		peak    int
	}
}

func NewGame() *Game {
	g := &Game{
		world: world.NewWorld(),
	}
	g.populationHistory.maxLen = 100
	return g
}

func (g *Game) Update() error {
	if g.tick%world.TicksPerFrame == 0 {
		g.world.Update()

		g.populationHistory.rabbits = append(g.populationHistory.rabbits, len(g.world.Rabbits))
		g.populationHistory.foxes = append(g.populationHistory.foxes, len(g.world.Foxes))

		if len(g.populationHistory.rabbits) > g.populationHistory.maxLen {
			g.populationHistory.rabbits = g.populationHistory.rabbits[1:]
			g.populationHistory.foxes = g.populationHistory.foxes[1:]
		}

		g.updatePopulationPeak()
	}
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	worldImage := ebiten.NewImage(world.MapWidth*graphics.TileSize, world.MapHeight*graphics.TileSize)
	graphics.RenderWorld(worldImage, g.world)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(worldImage, op)

	info := fmt.Sprintf("Population:\n\nRabbits: %d\nFoxes: %d",
		len(g.world.Rabbits),
		len(g.world.Foxes))

	text.Draw(screen,
		info,
		basicfont.Face7x13,
		world.MapWidth*graphics.TileSize+10,
		20,
		color.White)

	g.drawPopulationGraph(screen)
}

func (g *Game) updatePopulationPeak() {
	currentPeak := 0
	for i := range g.populationHistory.rabbits {
		if g.populationHistory.rabbits[i] > currentPeak {
			currentPeak = g.populationHistory.rabbits[i]
		}
		if g.populationHistory.foxes[i] > currentPeak {
			currentPeak = g.populationHistory.foxes[i]
		}
	}
	g.populationHistory.peak = currentPeak
}

func (g *Game) drawPopulationGraph(screen *ebiten.Image) {
	graphX := world.MapWidth*graphics.TileSize + 10
	graphY := 100
	minGraphHeight := 100
	graphWidth := 180

	graphHeight := minGraphHeight
	if g.populationHistory.peak > minGraphHeight {
		graphHeight = g.populationHistory.peak + 10
	}

	graphImg := ebiten.NewImage(graphWidth, graphHeight)
	graphImg.Fill(color.RGBA{20, 20, 20, 255})

	if len(g.populationHistory.rabbits) > 1 {
		for i := 1; i < len(g.populationHistory.rabbits); i++ {
			x1 := float64(graphWidth) * float64(i-1) / float64(g.populationHistory.maxLen)
			x2 := float64(graphWidth) * float64(i) / float64(g.populationHistory.maxLen)

			maxValue := float64(g.populationHistory.peak)
			if maxValue < float64(minGraphHeight) {
				maxValue = float64(minGraphHeight)
			}
			scaleY := float64(graphHeight) / maxValue

			y1 := float64(graphHeight) - (float64(g.populationHistory.rabbits[i-1]) * scaleY)
			y2 := float64(graphHeight) - (float64(g.populationHistory.rabbits[i]) * scaleY)
			ebitenutil.DrawLine(graphImg, x1, y1, x2, y2, color.White)

			y1 = float64(graphHeight) - (float64(g.populationHistory.foxes[i-1]) * scaleY)
			y2 = float64(graphHeight) - (float64(g.populationHistory.foxes[i]) * scaleY)
			ebitenutil.DrawLine(graphImg, x1, y1, x2, y2, color.RGBA{255, 0, 0, 255})
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(graphX), float64(graphY))
	screen.DrawImage(graphImg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return world.MapWidth*graphics.TileSize + 200, world.MapHeight * graphics.TileSize
}

func main() {
	rand.Seed(time.Now().UnixNano())

	graphics.LoadAssets()
	game := NewGame()

	// starting settings for the simulation are located in 'world.go'

	for y := 0; y < world.MapHeight; y++ {
		for x := 0; x < world.MapWidth; x++ {
			if rand.Float64() < world.InitialGrassDensity {
				game.world.Tiles[y][x].Grass = world.NewGrass()
			}
		}
	}

	for i := 0; i < world.InitialRabbitCount; i++ {
		x := rand.Intn(world.MapWidth)
		y := rand.Intn(world.MapHeight)
		game.world.Rabbits = append(game.world.Rabbits, world.NewRabbit(x, y))
	}

	for i := 0; i < world.InitialFoxCount; i++ {
		x := rand.Intn(world.MapWidth)
		y := rand.Intn(world.MapHeight)
		game.world.Foxes = append(game.world.Foxes, world.NewFox(x, y))
	}

	ebiten.SetWindowTitle("Rabbits and Foxes – Ecosystem")
	ebiten.SetWindowSize(956, 800) // fit for a 32x32 map

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
