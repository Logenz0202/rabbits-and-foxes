package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"rabbits-and-foxes/internal/graphics"
	"rabbits-and-foxes/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	World             *world.World
	Tick              int
	SimulationTick    int
	IsPaused          bool
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
	g.IsPaused = false
	g.PopulationHistory.MaxLen = 100
	g.CreateLife()
	return g
}

func (g *Game) TogglePause() {
	g.IsPaused = !g.IsPaused
}

func (g *Game) ToggleRestart() {
	g.World = world.NewWorld()
	g.Tick = 0
	g.SimulationTick = 0
	g.PopulationHistory.Rabbits = []int{}
	g.PopulationHistory.Foxes = []int{}
	g.PopulationHistory.Peak = 0
	g.CreateLife()
	g.IsPaused = false
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.TogglePause()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.ToggleRestart()
	}

	g.Tick++

	if !g.IsPaused && g.Tick%world.SimulationSpeed == 0 {
		g.SimulationTick++
		g.World.Update()

		g.PopulationHistory.Rabbits = append(g.PopulationHistory.Rabbits, len(g.World.Rabbits))
		g.PopulationHistory.Foxes = append(g.PopulationHistory.Foxes, len(g.World.Foxes))

		if len(g.PopulationHistory.Rabbits) > g.PopulationHistory.MaxLen {
			g.PopulationHistory.Rabbits = g.PopulationHistory.Rabbits[1:]
			g.PopulationHistory.Foxes = g.PopulationHistory.Foxes[1:]
		}
	}

	g.updatePopulationPeak()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	worldImage := ebiten.NewImage(world.MapWidth*graphics.TileSize, world.MapHeight*graphics.TileSize)
	graphics.RenderWorld(worldImage, g.World)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(worldImage, op)

	info := fmt.Sprintf("Population:\n\nRabbits: %d\nFoxes: %d\n",
		len(g.World.Rabbits),
		len(g.World.Foxes))

	buttonsText := "\n\nSPACE: pause/play\nR: restart"

	text.Draw(screen,
		info,
		basicfont.Face7x13,
		world.MapWidth*graphics.TileSize+10,
		20,
		color.White)

	text.Draw(screen,
		buttonsText,
		basicfont.Face7x13,
		world.MapWidth*graphics.TileSize+10,
		50,
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

func (g *Game) CreateLife() {
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
}
