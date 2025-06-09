package game

const (
	MapWidth  = 100
	MapHeight = 100
)

type Tile struct {
	Grass *Grass
}

type World struct {
	Tiles [][]Tile
}

func NewWorld() *World {
	tiles := make([][]Tile, MapHeight)
	for y := range tiles {
		tiles[y] = make([]Tile, MapWidth)
		for x := range tiles[y] {
			tiles[y][x] = Tile{Grass: nil}
		}
	}
	return &World{Tiles: tiles}
}

func (w *World) Update() {
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			if w.Tiles[y][x].Grass != nil {
				w.Tiles[y][x].Grass.Update(w, x, y)
			}
		}
	}
}
