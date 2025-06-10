package world

const (
	MapWidth  = 32
	MapHeight = 32
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
