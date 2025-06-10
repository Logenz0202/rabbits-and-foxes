package graphics

import (
	"rabbits-and-foxes/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const TileSize = 32
const minAlpha = 0.5

func RenderWorld(screen *ebiten.Image, w *world.World) {
	for y := 0; y < world.MapHeight; y++ {
		for x := 0; x < world.MapWidth; x++ {
			tile := w.Tiles[y][x]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			if tile.Grass != nil && tile.Grass.Amount > 0.1 {
				alpha := minAlpha + (tile.Grass.Amount/world.MaxGrassAmount)*(1.0-minAlpha)
				op.ColorM.Scale(1, 1, 1, alpha)
				screen.DrawImage(GrassTile, op)
			} else {
				screen.DrawImage(DirtTile, op)
			}
		}
	}

	for _, r := range w.Rabbits {
		if r.IsAlive() {
			op := &ebiten.DrawImageOptions{}
			if r.Direction < 0 {
				op.GeoM.Scale(-1, 1) // odbicie w poziomie
			}
			op.GeoM.Translate(float64(r.X*TileSize), float64(r.Y*TileSize))
			screen.DrawImage(Rabbit, op)
		}
	}
}
