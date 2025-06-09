package graphics

import (
	"rabbits-and-foxes/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const TileSize = 32

const minAlpha = 0.5

func RenderWorld(screen *ebiten.Image, world *game.World) {
	for y := 0; y < game.MapHeight; y++ {
		for x := 0; x < game.MapWidth; x++ {
			tile := world.Tiles[y][x]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))

			if tile.Grass != nil && tile.Grass.Amount > 0.1 {
				alpha := minAlpha + (tile.Grass.Amount/game.MaxGrassAmount)*(1.0-minAlpha)
				op.ColorM.Scale(1, 1, 1, alpha)
				screen.DrawImage(GrassTile, op)
			} else {
				screen.DrawImage(DirtTile, op)
			}
		}
	}
}
