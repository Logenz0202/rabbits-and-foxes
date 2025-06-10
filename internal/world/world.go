package world

const (
	MapWidth  = 16
	MapHeight = 16
)

type Tile struct {
	Grass *Grass
}

type World struct {
	Tiles   [][]Tile
	Rabbits []*Rabbit
}

func NewWorld() *World {
	tiles := make([][]Tile, MapHeight)
	for y := range tiles {
		tiles[y] = make([]Tile, MapWidth)
		for x := range tiles[y] {
			tiles[y][x] = Tile{Grass: nil}
		}
	}
	return &World{
		Tiles:   tiles,
		Rabbits: []*Rabbit{},
	}
}

func (w *World) Update() {
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			if w.Tiles[y][x].Grass != nil {
				w.Tiles[y][x].Grass.Update(w, x, y)
			}
		}
	}

	var newRabbits []*Rabbit
	var aliveRabbits []*Rabbit

	for _, r := range w.Rabbits {
		if !r.IsAlive() {
			continue
		}

		r.Move(w)
		r.Eat(w)

		r.Energy -= RabbitEnergyLossPerTick
		if r.ReproductionCooldown > 0 {
			r.ReproductionCooldown--
		}

		if r.CanReproduce() {
			for _, other := range w.Rabbits {
				if other != r && other.CanReproduce() &&
					Abs(r.X-other.X) <= 1 && Abs(r.Y-other.Y) <= 1 {
					child := NewRabbit(r.X, r.Y)
					newRabbits = append(newRabbits, child)
					r.ReproductionCooldown = RabbitReproductionCD
					other.ReproductionCooldown = RabbitReproductionCD
					break
				}
			}
		}

		aliveRabbits = append(aliveRabbits, r)
	}

	w.Rabbits = append(aliveRabbits, newRabbits...)
}

func (w *World) IsOccupiedByRabbit(x, y int) *Rabbit {
	for _, rabbit := range w.Rabbits {
		if rabbit.IsAlive() && rabbit.X == x && rabbit.Y == y {
			return rabbit
		}
	}
	return nil
}
