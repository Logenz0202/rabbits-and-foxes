package game

import (
	"fmt"
	"image/color"
	"rabbits-and-foxes/internal/graphics"
	"rabbits-and-foxes/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	World             *world.World
	Tick              int
	PopulationHistory struct {
		Rabbits []int
		Foxes   []int
		MaxLen  int
		Peak    int
	}
}

func NewGame() *Game {
	g := &Game{
		World: world.NewWorld(),
	}
	g.PopulationHistory.MaxLen = 100
	return g
}

func (g *Game) Update() error {
	if g.Tick%world.TicksPerFrame == 0 {
		g.World.Update()

		g.PopulationHistory.Rabbits = append(g.PopulationHistory.Rabbits, len(g.World.Rabbits))
		g.PopulationHistory.Foxes = append(g.PopulationHistory.Foxes, len(g.World.Foxes))

		if len(g.PopulationHistory.Rabbits) > g.PopulationHistory.MaxLen {
			g.PopulationHistory.Rabbits = g.PopulationHistory.Rabbits[1:]
			g.PopulationHistory.Foxes = g.PopulationHistory.Foxes[1:]
		}

		g.updatePopulationPeak()
	}
	g.Tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	worldImage := ebiten.NewImage(world.MapWidth*graphics.TileSize, world.MapHeight*graphics.TileSize)
	graphics.RenderWorld(worldImage, g.World)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(worldImage, op)

	info := fmt.Sprintf("Population:\n\nRabbits: %d\nFoxes: %d",
		len(g.World.Rabbits),
		len(g.World.Foxes))

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
	for i := range g.PopulationHistory.Rabbits {
		if g.PopulationHistory.Rabbits[i] > currentPeak {
			currentPeak = g.PopulationHistory.Rabbits[i]
		}
		if g.PopulationHistory.Foxes[i] > currentPeak {
			currentPeak = g.PopulationHistory.Foxes[i]
		}
	}
	g.PopulationHistory.Peak = currentPeak
}

func (g *Game) drawPopulationGraph(screen *ebiten.Image) {
	graphX := world.MapWidth*graphics.TileSize + 10
	graphY := 100
	minGraphHeight := 100
	graphWidth := 180

	graphHeight := minGraphHeight
	if g.PopulationHistory.Peak > minGraphHeight {
		graphHeight = g.PopulationHistory.Peak + 10
	}

	graphImg := ebiten.NewImage(graphWidth, graphHeight)
	graphImg.Fill(color.RGBA{20, 20, 20, 255})

	if len(g.PopulationHistory.Rabbits) > 1 {
		for i := 1; i < len(g.PopulationHistory.Rabbits); i++ {
			x1 := float64(graphWidth) * float64(i-1) / float64(g.PopulationHistory.MaxLen)
			x2 := float64(graphWidth) * float64(i) / float64(g.PopulationHistory.MaxLen)

			maxValue := float64(g.PopulationHistory.Peak)
			if maxValue < float64(minGraphHeight) {
				maxValue = float64(minGraphHeight)
			}
			scaleY := float64(graphHeight) / maxValue

			y1 := float64(graphHeight) - (float64(g.PopulationHistory.Rabbits[i-1]) * scaleY)
			y2 := float64(graphHeight) - (float64(g.PopulationHistory.Rabbits[i]) * scaleY)
			ebitenutil.DrawLine(graphImg, x1, y1, x2, y2, color.White)

			y1 = float64(graphHeight) - (float64(g.PopulationHistory.Foxes[i-1]) * scaleY)
			y2 = float64(graphHeight) - (float64(g.PopulationHistory.Foxes[i]) * scaleY)
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
