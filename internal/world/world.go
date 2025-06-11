package world

const (
	MapWidth            = 32   // number of vertical tiles
	MapHeight           = 32   // number of horizontal tiles
	InitialGrassDensity = 0.15 // % of tiles with pre-planted grass
	InitialRabbitCount  = 100  // initial number of rabbits
	InitialFoxCount     = 20   // initial number of foxes
	TicksPerFrame       = 5    // number of world updates per frame
)

type Tile struct {
	Grass *Grass
}

type World struct {
	Tiles   [][]Tile
	Rabbits []*Rabbit
	Foxes   []*Fox
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
		Foxes:   []*Fox{},
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

		if r.CanReproduce() {
			for _, other := range w.Rabbits {
				if other != r && other.IsAlive() && other.CanReproduce() &&
					IsAdjacent(r.X, r.Y, other.X, other.Y) {
					child := NewRabbit(r.X, r.Y)
					newRabbits = append(newRabbits, child)
					r.ReproductionCooldown = RabbitReproductionCD
					other.ReproductionCooldown = RabbitReproductionCD
					break
				}
			}
		}

		r.Energy -= RabbitEnergyLossPerTick
		if r.ReproductionCooldown > 0 {
			r.ReproductionCooldown--
		}

		aliveRabbits = append(aliveRabbits, r)
	}

	w.Rabbits = append(aliveRabbits, newRabbits...)

	var newFoxes []*Fox
	var aliveFoxes []*Fox

	for _, f := range w.Foxes {
		if !f.IsAlive() {
			continue
		}

		f.Move(w)

		if f.CanReproduce() {
			for _, other := range w.Foxes {
				if other != f && other.IsAlive() && other.CanReproduce() &&
					IsAdjacent(f.X, f.Y, other.X, other.Y) {
					child := NewFox(f.X, f.Y)
					newFoxes = append(newFoxes, child)
					f.ReproductionCooldown = FoxReproductionCD
					other.ReproductionCooldown = FoxReproductionCD
					break
				}
			}
		}

		f.Energy -= FoxEnergyLossPerTick
		if f.ReproductionCooldown > 0 {
			f.ReproductionCooldown--
		}

		aliveFoxes = append(aliveFoxes, f)
	}

	w.Foxes = append(aliveFoxes, newFoxes...)
}

func (w *World) IsOccupiedByRabbit(x, y int) *Rabbit {
	for _, rabbit := range w.Rabbits {
		if rabbit.IsAlive() && rabbit.X == x && rabbit.Y == y {
			return rabbit
		}
	}
	return nil
}

func (w *World) IsOccupiedByFox(x, y int) *Fox {
	for _, fox := range w.Foxes {
		if fox.IsAlive() && fox.X == x && fox.Y == y {
			return fox
		}
	}
	return nil
}
